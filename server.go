package main


import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/host"
	"os"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {		
		host,_ := host.Info()
		
		c.JSON(http.StatusOK, gin.H {
			"Hostname": host.Hostname,
			"Uptime": host.Uptime,
			"Procs": host.Procs,
			"OS": host.OS,
			"Platform": host.Platform,
			"PlatformFamily": host.PlatformFamily,
			"PlatformVersion": host.PlatformVersion,
			"VirtualizationSystem": host.VirtualizationSystem,
			"VirtualizationRole": host.VirtualizationRole,
		})
	})

	r.GET("/env", func(c *gin.Context) {	

		env := os.Environ()

		c.JSON(http.StatusOK, gin.H {
			"env": env,			
		})
	})

	r.GET("/ready", func(c *gin.Context) {	

		host,_ := host.Info()

		if host.Uptime < 30 {
        	c.JSON(http.StatusServiceUnavailable, gin.H {
				"message": "Not Ready :(",			
			})	
    	} else {
    		c.JSON(http.StatusOK, gin.H {
				"message": "Ready :)",			
			})	
    	}
		
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
