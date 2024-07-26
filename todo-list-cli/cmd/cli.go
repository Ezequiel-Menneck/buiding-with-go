package cmd

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/spf13/cobra"
	"time"
	"todo-list-cli/utils"
)

var activityName string
var category string
var timer int
var activities = make(map[string]*ActivityType)

type ActivityType struct {
	ActivityName string
	Category     string
	TimeToDone   int
	EndTime      time.Time
	UpdateChan   chan time.Duration
}

var addActivity = &cobra.Command{
	Use:     "addActivity",
	Short:   "A brief description of your application",
	Aliases: []string{"add"},
	Run: func(cmd *cobra.Command, args []string) {
		if _, ok := activities[activityName]; ok {
			fmt.Println(activityName, "exists in your activities map")
			return
		}

		fmt.Printf("Adding activity: %v for your list\n", activityName)

		futureTime := time.Now().Add(time.Duration(timer) * time.Minute)
		updateChan := make(chan time.Duration)
		activity := &ActivityType{
			ActivityName: activityName,
			Category:     category,
			TimeToDone:   timer,
			EndTime:      futureTime,
			UpdateChan:   updateChan,
		}
		activities[activityName] = activity

		go func(t *ActivityType) {
			newTimer := time.NewTimer(time.Duration(timer) * time.Second)
			defer newTimer.Stop()
			for {
				select {
				case newDuration := <-t.UpdateChan:
					newTimer.Stop()
					t.EndTime = time.Now().Add(newDuration)
					newTimer = time.NewTimer(time.Duration(timer) * time.Second)
				case <-newTimer.C:
					newTimer.Stop()
					delete(activities, t.ActivityName)
					err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
					if err != nil {
						panic(err)
					}
					err = beeep.Notify(activityName, "Its time to do your activity: "+t.ActivityName, "")
					if err != nil {
						panic(err)
					}
				}
			}

		}(activity)
	},
}

var updateActivity = &cobra.Command{
	Use:   "updateActivity",
	Short: "A brief description of your application",
	Run: func(cmd *cobra.Command, args []string) {
		task, exists := activities[activityName]
		if exists {
			task.UpdateChan <- time.Duration(timer) * time.Minute
			fmt.Printf("Activity %v updated with new duration: %v\n", activityName, category)
		} else {
			fmt.Printf("Activity %v not found.\n", activityName)
		}
	},
}

var getAllActivities = &cobra.Command{
	Use:   "getAllActivities",
	Short: "A brief description of your application",
	Run: func(cmd *cobra.Command, args []string) {
		if len(activities) == 0 {
			fmt.Println("You dont have any activity for today :D")
			return
		}

		fmt.Println("Your activities list for the day is: ")
		now := time.Now()
		i := 1
		for _, v := range activities {
			durationStr := utils.FormatDate(v.EndTime.Sub(now).String())
			fmt.Printf("\t %v. Activity: %v in %v \n", i, v.ActivityName, durationStr)
			i++
		}
	},
}

var getActivityByName = &cobra.Command{
	Use:   "getActivityByName",
	Short: "A brief description of your application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This activity is: ")

		now := time.Now()
		myActivity := activities[activityName]
		durationStr := utils.FormatDate(myActivity.EndTime.Sub(now).String())
		fmt.Printf("\t Activity: %v in %v \n", myActivity.ActivityName, durationStr)

	},
}

var getActivitiesByCategory = &cobra.Command{
	Use:   "getActivitiesByCategory",
	Short: "A brief description of your application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("The list of activities of category %v is: \n", category)

		activitiesByCategory := make(map[string]ActivityType)
		for _, v := range activities {
			if v.Category == category {
				activitiesByCategory[v.ActivityName] = *v
			}
		}

		i := 1
		for _, v := range activitiesByCategory {
			fmt.Printf("\t %v - Activity: %v, of Category: %v\n", i, v.ActivityName, v.Category)
			i++
		}

	},
}

func init() {
	rootCmd.AddCommand(addActivity)
	addActivity.Flags().StringVarP(&activityName, "activity", "n", "none", "What activity you need to be remembered?")
	addActivity.Flags().StringVarP(&category, "category", "c", "", "What category you wanna to add for your activity?")
	addActivity.Flags().IntVarP(&timer, "time", "t", 10, "Time to call notification to remember the activity")

	rootCmd.AddCommand(updateActivity)
	updateActivity.Flags().StringVarP(&activityName, "activity", "n", "", "What activity you need to be remembered?")
	updateActivity.Flags().StringVarP(&category, "category", "c", "", "What category you wanna to add for your activity?")
	updateActivity.Flags().IntVarP(&timer, "time", "t", 10, "Time to call notification to remember the activity")

	rootCmd.AddCommand(getAllActivities)

	rootCmd.AddCommand(getActivityByName)
	getActivityByName.Flags().StringVarP(&activityName, "activity", "c", "", "What activity you need to be remembered?")

	rootCmd.AddCommand(getActivitiesByCategory)
	getActivitiesByCategory.Flags().StringVarP(&category, "category", "c", "", "What category you need to be remembered?")

}
