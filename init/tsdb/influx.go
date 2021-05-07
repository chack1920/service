package tsdb

import (
	"context"

	client "github.com/influxdata/influxdb1-client/v2"
)

// Influx influx存储
type Influx struct {
	cli client.Client
}

func NewInflux(cli client.Client) TSDB {
	a := new(Influx)
	a.cli = cli
	return a
}

func (a *Influx) Write(ctx context.Context, database string, rows []Row) (err error) {
	ps := make([]*client.Point, 0)
	for _, row := range rows {
		pt, err := client.NewPoint(
			row.TableName,
			row.Tags,
			row.Fields,
			row.Ts,
		)
		if err != nil {
			return err
		}
		ps = append(ps, pt)
	}

	if len(ps) > 0 {
		bp, err := client.NewBatchPoints(client.BatchPointsConfig{
			Database:        database,
			Precision:       "ms",
			RetentionPolicy: "autogen",
		})
		if err != nil {
			return err
		}
		bp.AddPoints(ps)
		return a.cli.Write(bp)
	}

	return nil
}

func (a *Influx) Query(ctx context.Context, database string, query string) (res []client.Result, err error) {
	q := client.Query{
		Command:  query,
		Database: database,
	}

	if response, err := a.cli.Query(q); err == nil {
		if response.Error() != nil {
			return nil, response.Error()
		}
		return response.Results, nil
	} else {
		return nil, err
	}
}

func (a *Influx) QueryFilter(ctx context.Context, database string, query []map[string]interface{}) (res []client.Result, err error) {
	return nil, nil
}
