package help

import (
	"fmt"
	"os"

	"skeely/command/version"
)

const USAGE = `skeely %s Is a Schema Linter for MySQL, this tool help to identifying
some common and uncommon mistakes on data model.

USAGE:
	skeely [--help | --path | --version]

OPTIONS:
  --comment               Send summary as comment into GitHub.
  --help                  Show this help.
  --path                  Path of the directory to start to find *.sql to lint.
  --git                   Auto identifying git changed files, require --path option.
  --github-pull-request   Pull request number.
  --github-repository     Repository path on github.
  --github-token          Token to auth in github.
  --version               Print version numbers.

EXAMPLES:

  # Lint directory
  $ skeely --path=assets/examples/

  # Lint specific file
  $ skeely --path=assets/examples/case01.sql

  # Lint and push summary as comment into GitHub Pull Request.
  $ skeely --path=assets/examples/case01.sql \
           --github-comment \
           --github-token=${{github.token}} \
           --github-repository=$GITHUB_REPOSITORY \
           --github-pull-request=${{github.event.pull_request.number}}

For more help, plese visit: https://github.com/debeando/skeely
`

func Show(rc int) {
	fmt.Printf(USAGE, version.VERSION)
	os.Exit(rc)
}