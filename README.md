# Stat Service

A go micro-service for calculating a least squares fit.

## Getting Started

### Prerequisites

- make
- go 1.24 or later
- Docker

### Installation

```bash
git clone https://github.com/octalope/stat-service.git
cd stat-service
make all
```

### Running the Service

```bash
make run
```

to run the service locally, or

```bash
make run-docker
```

to run the service in the Docker container

## API Endpoints

| Method | Endpoint         | Description                           |
|--------|------------------|---------------------------------------|
| GET    | `/health`        | Health check                          |
| POST   | `/lsf`           | Submit new data for least squares fit |

### Example Requests

#### Get `/health`

```bash
curl http://localhost:8080/health
```

will always respond with `HTTP/1.1 200 OK` and the text body `OK`.

#### POST Data to `/lsf`

This returns a least-squares-fit for column 1 (y) versus column 0 (x) of the data.

```bash
curl -X POST http://localhost:8080/lsf \
  -H "Content-Type: application/json" \
  -d '{ 
        "data": [
          [1, 2.85,   2],
          [2, 5.10,   4],
          [3, 6.9,    6],
          [4, 9.1,    8],
          [5, 10.9,  10],
          [6, 12.85, 12],
          [7, 15.2,  14]
        ], 
        "x_col": 0,
        "y_col": 1 
      }'
```

where:

| Field   | Description                                                      |
|---------|------------------------------------------------------------------|
| `data`  | Two-dimensional array of floating point data                     |
| `x_col` | Index of the independent variable for the least-squares fit      |
| `y_col` | Index of the dependent variable for the least-squares fit        |

the JSON response will be

```json
{
  "m":2.019642857142857,
  "dm":0.05676212448023936,
  "b":0.9071428571428584,
  "db":0.25384793777024023,
  "rSquared":0.9990136079923525
}  
```

where:

| Field     | Description                          |
|-----------|--------------------------------------|
| `m`       | Slope                                |
| `dm`      | Uncertainty in the slope             |
| `b`       | Y-intercept                          |
| `db`      | Uncertainty of the y-intercept       |
| `rSquared`| Correlation coefficient              |

## License

[MIT](./LICENSE.md)
