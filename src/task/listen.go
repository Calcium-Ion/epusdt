package task

import "github.com/robfig/cron/v3"

func Start() {
	c := cron.New()

	// Create job instances
	usdtJob := UsdtRateJob{}
	trc20Job := ListenTrc20Job{}

	// Run USDT rate job immediately
	usdtJob.Run()

	// Schedule jobs
	c.AddJob("@every 60s", usdtJob)
	c.AddJob("@every 5s", trc20Job)
	c.Start()
}
