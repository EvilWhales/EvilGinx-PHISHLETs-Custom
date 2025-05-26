	func showAd() {
		lred := color.New(color.FgHiRed)
		lyellow := color.New(color.FgHiYellow)
		white := color.New(color.FgHiWhite)
		message := fmt.Sprintf("%s: %s %s", lred.Sprint("Evilginx Mastery Course"), lyellow.Sprint("https://academy.breakdev.org/evilginx-mastery"), white.Sprint("(learn how to create phishlets)"))
		log.Info("%s", message)
	}

	var google_bypass = flag.Bool("google-bypass", false, "Enable Google Bypass")

	func init() {
		flag.Parse()
		if *google_bypass {
			// Ensure the DISPLAY environment variable is set
			display := ":99"
			if disp := getenv("DISPLAY", ""); disp != "" {
				display = disp
			}

			// Kill any existing Chrome instances with remote debugging on port 9222
			exec.Command("pkill", "-f", "google-chrome.*--remote-debugging-port=9222").Run()
			log.Debug("Killed all google-chrome instances running in debug mode on port 9222")

			// Start google-chrome in debug mode
			cmd := exec.Command("google-chrome", "--remote-debugging-port=9222", "--no-sandbox")

			// Capture standard output and error
			var out bytes.Buffer
			var stderr bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &stderr

			// Set environment variables if necessary (e.g., DISPLAY)
			cmd.Env = append(cmd.Env, fmt.Sprintf("DISPLAY=%s", display))

			err := cmd.Start()
			if err != nil {
				log.Error("Failed to start google-chrome in debug mode: %v", err)
				log.Error("Command output: %s", stderr.String())
				return
			}
			log.Debug("Started google-chrome in debug mode on port 9222")

			// Optionally wait for the command to finish and capture its output
			go func() {
				err = cmd.Wait()
				if err != nil {
					log.Error("google-chrome process exited with error: %v", err)
					log.Error("Command output: %s", stderr.String())
				}
			}()
			// Ensure a browser instance is available
			launcher.NewBrowser().MustGet()
		}
	}

	func getenv(key, fallback string) string {
		value := os.Getenv(key)
		if value == "" {
			return fallback
		}
		return value
	}