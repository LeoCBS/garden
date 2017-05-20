# Garder server

Receive garden information from arduino sensors

# API

## Parameters [/parameters]

### List All Parameters [GET]

+ Response 200 (application/json)

        [
            {
                "name": "humidity",
                "value": "80",
                "measure": "percent",
                "lastUpdate": "2015-08-05T08:40:51.620Z"
            }
        ]
        

### Create a New Parameter [POST]

You may create your own parameter using this action. It takes a JSON
object containing a parameters values.

+ Request (application/json)

        {
            "name": "humidity",
            "value": "80",
            "measure": "percent",
        }

+ Response 201 (application/json)

    + Headers

            Location: /parameter/UUID

    + Body

            {
                "created": "2015-08-05T08:40:51.620Z"
            }
