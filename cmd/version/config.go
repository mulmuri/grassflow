package version

import "fmt"

var (
	WorkflowName = "version"
	version = "1.0.0"
)



func Run() {
	output := fmt.Sprintf("Grassflow: %s", version)
	fmt.Println(output)
}