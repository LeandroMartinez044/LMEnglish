package dependencies

import (
	"lmenglish/internal/api/handlers/recordhdl"
	"lmenglish/internal/core/ports"
	"lmenglish/internal/core/services/recordsrv"
)

type Definition struct {

	//
	// Repositories
	//

	//
	// Services
	//
	RecordService ports.RecordService

	//
	// Handler
	//
	RecordHandler *recordhdl.Handler
}

func Init() Definition {
	d := Definition{}

	d.RecordService = recordsrv.New()
	d.RecordHandler = recordhdl.New(d.RecordService)

	return d
}
