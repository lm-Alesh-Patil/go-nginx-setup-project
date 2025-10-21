# go-nginx-setup-project
Let’s build this step by step like a real mini project: We’ll make a Go app, then put NGINX in front of it inside Docker — all running locally.

go-nginx-setup-project/
├── docker-compose.yml
├── nginx/
│   └── nginx.conf
└── app/
    ├── main.go
    └── Dockerfile

# for running all containers in docker 
docker compose up --build

# for clean up container (Removal)
docker compose down -v

# for running autocannon
docker compose run autocannon

12 packages are looking for funding
  run `npm fund` for details
npm notice
npm notice New major version of npm available! 10.8.2 -> 11.6.2
npm notice Changelog: https://github.com/npm/cli/releases/tag/v11.6.2
npm notice To update run: npm install -g npm@11.6.2
npm notice
Running 20s test @ http://nginx
50 connections


┌─────────┬───────┬───────┬───────┬───────┬──────────┬──────────┬────────┐
│ Stat    │ 2.5%  │ 50%   │ 97.5% │ 99%   │ Avg      │ Stdev    │ Max    │
├─────────┼───────┼───────┼───────┼───────┼──────────┼──────────┼────────┤
│ Latency │ 13 ms │ 28 ms │ 63 ms │ 79 ms │ 30.15 ms │ 12.49 ms │ 182 ms │
└─────────┴───────┴───────┴───────┴───────┴──────────┴──────────┴────────┘
┌───────────┬────────┬────────┬────────┬────────┬─────────┬─────────┬────────┐
│ Stat      │ 1%     │ 2.5%   │ 50%    │ 97.5%  │ Avg     │ Stdev   │ Min    │
├───────────┼────────┼────────┼────────┼────────┼─────────┼─────────┼────────┤
│ Req/Sec   │ 1,081  │ 1,081  │ 1,702  │ 1,923  │ 1,636.1 │ 242.25  │ 1,081  │
├───────────┼────────┼────────┼────────┼────────┼─────────┼─────────┼────────┤
│ Bytes/Sec │ 215 kB │ 215 kB │ 339 kB │ 383 kB │ 326 kB  │ 48.2 kB │ 215 kB │
└───────────┴────────┴────────┴────────┴────────┴─────────┴─────────┴────────┘

Req/Bytes counts sampled once per second.
# of samples: 20

33k requests in 20.18s, 6.51 MB read

# Summary
Over 20 seconds:
33,000 requests successfully completed
6.51 MB of total response data transferred

# What You Just Proved
Your NGINX reverse proxy + Go services setup can handle ~1.6k requests/second easily.
NGINX evenly distributed requests across 3 Go apps (load balanced).
You can scale further by adding goapp4, goapp5, etc., to the upstream in nginx.conf.