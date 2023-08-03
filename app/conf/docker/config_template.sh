cat <<EOF >conf/config.json
{
  "app": {
    "name": "go-ddd-template",
    "env": "prod",
    "port": 8001,
    "debug_port": 18001,
    "deploy_mode": "${YS_START_MODEL}"
  },
  "logger": {
    "log_level": "info",
    "max_backups": 3,
    "max_size": 100
  },
  "mysql_rw": {
    "host": "${YS_MYSQL_HOST}",
    "port": ${YS_MYSQL_PORT},
    "user": "${YS_MYSQL_USER}",
    "password": "${YS_MYSQL_PWD}",
    "db": "${YS_CONSOLE_DB}",
    "charset": "utf8mb4",
    "max_open_count": 32,
    "max_idle_count": 32,
    "log": {
      "max_backups": 3,
      "max_size": 100
    }
  },
  "redis": {
    "host": "${YS_REDIS_HOST}",
    "port": ${YS_REDIS_PORT},
    "auth": "${YS_REDIS_PWD}",
    "pool_size": 128
  }
}
EOF