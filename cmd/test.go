package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

var amount int

func init() {
	testCmd.Flags().IntVarP(&amount, "amount", "q", 2, "amount of games to run in parallel")

	rootCmd.AddCommand(testCmd)
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "run tests in parallel",
	Long: `test runs rcss games in parallel and output their results to stdout

Must have two teams in the current folder, inside folders named "team" and "gliders"`,

	RunE: func(cmd *cobra.Command, args []string) error {
		if amount <= 0 {
			return errors.New("amount must be greater then 0")
		}

		if _, err := os.Stat("./team/start.sh"); err != nil {
			return err
		}

		if _, err := os.Stat("./gliders/start.sh"); err != nil {
			return err
		}

		var wg sync.WaitGroup
		port := 5000

		for i := 1; i <= amount; i++ {
			wg.Add(1)

			go func(p int, i int) {
				ifVerbose(fmt.Sprint("Spawning game", i))
				defer wg.Done()

				filename := fmt.Sprintf("teste_%d.csv", i)

				teamCmd := exec.Command("./team/start.sh", "-t", "team")
				teamCmd.Args = append(teamCmd.Args, "-p", fmt.Sprint(p))
				teamCmd.Args = append(teamCmd.Args, "-P", fmt.Sprint(p+1))

				glidersCmd := exec.Command("./gliders/start.sh", "-t", "gliders")
				glidersCmd.Args = append(glidersCmd.Args, "-p", fmt.Sprint(p))
				glidersCmd.Args = append(glidersCmd.Args, "-P", fmt.Sprint(p+1))

				// server::synch_mode=true CSVSaver::save=true CSVSaver::filename=$CSV_FILENAME server::auto_mode=true server::port=$1 server::coach_port=$2 server::olcoach_port=$2
				serverCmd := exec.Command("rcssserver", "server::synch_mode=true", "server::auto_mode=true", "CSVSaver::save=true", "server::penalty_shoot_outs=false", "server::nr_extra_halfs=0")
				serverCmd.Args = append(serverCmd.Args, fmt.Sprintf("server::port=%d", p))
				serverCmd.Args = append(serverCmd.Args, fmt.Sprintf("server::coach_port=%d", p+1))
				serverCmd.Args = append(serverCmd.Args, fmt.Sprintf("server::olcoach_port=%d", p+2))
				serverCmd.Args = append(serverCmd.Args, fmt.Sprintf("CSVSaver::filename=%s", filename))

				serverCmd.Stdout, _ = os.Create(fmt.Sprint("server_stdout_", i))
				serverCmd.Stderr, _ = os.Create(fmt.Sprint("server_stderr_", i))

				teamCmd.Stdout, _ = os.Create(fmt.Sprint("team_stdout_", i))
				teamCmd.Stderr, _ = os.Create(fmt.Sprint("team_stderr_", i))

				glidersCmd.Stdout, _ = os.Create(fmt.Sprint("gliders_stdout_", i))
				glidersCmd.Stderr, _ = os.Create(fmt.Sprint("gliders_stderr_", i))

				teamCmd.Start()
				glidersCmd.Start()

				if err := serverCmd.Run(); err != nil {
					ifVerbose(fmt.Sprint(i, "server", err))
				}

				if r, err := ioutil.ReadFile(filename); err == nil {
					fmt.Println(strings.Split(string(r), "\n")[1])
				}

				os.Remove(filename)
			}(port, i)

			port += 3
		}

		wg.Wait()

		return nil
	},
}
