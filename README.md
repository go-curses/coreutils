# Go-Curses Core Utilities

## github.com/go-curses/coreutils/errors

Helpers for working with Go errors.

## github.com/go-curses/coreutils/path

Simple file path related utilities for reducing code duplication.

## github.com/go-curses/coreutils/replace

Utilities for searching and replacing content within files.

### github.com/go-curses/coreutils/replace/cmd/rpl

```
> go install github.com/go-curses/coreutils/replace/cmd/rpl
> $GOPATH/bin/rpl --help
NAME:
   rpl - command line search and replace

USAGE:
   rpl [global options] <search> <replace> <path> [...paths]

VERSION:
   0.1.0

DESCRIPTION:
   Search and replace content within files using strings or regular expressions.

GLOBAL OPTIONS:
   --regex, -p                         search and replace arguments are regular expressions (default: false)
   --recurse, -R                       recurse into sub-directories (default: false)
   --dry-run, -n                       report what would have otherwise been done (default: false)
   --all, -a                           include files and directories that start with a "." (default: false)
   --ignore-case, -i                   perform a case-insensitive search (default: false)
   --quiet, -q                         run silently, ignored if --dry-run is also used (default: false)
   --backup, -b                        make backups before replacing content (default: false)
   --backup-extension value, -B value  specify the backup file extension to use (default: "bak")
   --show-diff, -D                     include unified diffs of all changes in the output (default: false)
   --tmpdir value, -T value            specify the tmpdir to use (default: "/tmp") [$TMPDIR]
   --help, -h                          show help (default: false)
   --version, -v                       print the version (default: false)
```
