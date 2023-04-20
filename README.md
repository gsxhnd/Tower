# Tower

```shell
docker build -f Dockerfile -t mosquitto:dev .
docker run --name mosquitto -p 1883:1883 -p 9001:9001 mosquitto:dev
```

TOPIC `$CONTROL/dynamic-security/v1/response`
$CONTROL/dynamic-security/v1

```json
{
  "commands": [
    {
      "command": "listClients",
      "verbose": false,
      "count": -1,
      "offset": 0
    }
  ]
}
```

```json
{
  "commands": [
    {
      "command": "listClients",
      "verbose": true,
      "count": 1,
      "offset": 0,
      "username": "admin"
    }
  ]
}
```

```json
{
  "commands": [
    {
      "command": "createClient",
      "username": "admin01",
      "password": "000000",
      "clientid": "",
      "textname": "Admin user",
      "textdescription": "",
      "roles": [
        {
          "rolename": "admin"
        }
      ]
    }
  ]
}
```

```json
{
  "commands": [
    {
      "command": "listClients",
      "verbose": true,
      "count": -1,
      "offset": 0,
      "connected": true
    }
  ]
}
```

```json
{
  "commands": [
    {
      "command": "enableClient",
      "username": "admin01"
    }
  ]
}
```
