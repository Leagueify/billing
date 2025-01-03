# Billing Service

The Leagueify billing service is responsible for managing payments and financial
reporting. The Leagueify billing service uses [Go][go-download] using version
`1.23.0`.

## Key Functions

- Process payments for player registrations.
- Tracking billing history and provide financial reports for organizations and users.

## Development

### Prerequisites

- [Docker][docker-download] is installed and running.

### Getting Started

Local development of the Leagueify billing service uses docker for a consistent
development environment. Running the Leagueify billing service locally can be
accomplished using commands found in the [Makefile][repo-makefile]. During local
development changes will trigger a live reload, eliminating the requirement to
restart the docker image manually. This is accomplished by using the wonderful
tool [air][github-air]. The most common commands have been outlined below:

```bash
# start leagueify docker image
make start

# stop leagueify docker image removing attached volumes
make clean
```

The Leagueify billing service is ready for development once the banner output
is visible within the terminal. By default the Leagueify billing service api
docs are accessible at [http://localhost:6507/billing/docs][service-url]. The
banner below was created using the
[Text to ASCII Art Generator by Patorjk][patorjk-taag].

```
leagueify-billing-1  |
leagueify-billing-1  | '||'      '||''''|      |      ..|'''.|  '||'  '|' '||''''|  '||' '||''''| '||' '|'
leagueify-billing-1  |  ||        ||  .       |||    .|'     '   ||    |   ||  .     ||   ||  .     || |
leagueify-billing-1  |  ||        ||''|      |  ||   ||    ....  ||    |   ||''|     ||   ||''|      ||
leagueify-billing-1  |  ||        ||        .''''|.  '|.    ||   ||    |   ||        ||   ||         ||
leagueify-billing-1  | .||.....| .||.....| .|.  .||.  ''|...'|    '|..'   .||.....| .||. .||.       .||.
leagueify-billing-1  |
leagueify-billing-1  | '||''|.   '||' '||'      '||'      '||' '|.   '|'  ..|'''.|
leagueify-billing-1  |  ||   ||   ||   ||        ||        ||   |'|   |  .|'     '
leagueify-billing-1  |  ||'''|.   ||   ||        ||        ||   | '|. |  ||    ....
leagueify-billing-1  |  ||    ||  ||   ||        ||        ||   |   |||  '|.    ||
leagueify-billing-1  | .||...|'  .||. .||.....| .||.....| .||. .|.   '|   ''|...'|
leagueify-billing-1  |
```

[docker-download]: https://www.docker.com/get-started/
[github-air]: https://github.com/air-verse/air
[go-download]: https://go.dev/dl/
[patorjk-taag]: https://patorjk.com/software/taag/#p=display&f=Kban&t=LEAGUEIFY%0ABILLING
[repo-makefile]: ./Makefile
[service-url]: http://localhost:6507/billing/docs
