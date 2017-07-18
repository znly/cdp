// Code generated by cdpgen. DO NOT EDIT.

// Package memory implements the Memory domain.
package memory

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the Memory domain.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the Memory domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// GetDOMCounters invokes the Memory method.
func (d *domainClient) GetDOMCounters(ctx context.Context) (reply *GetDOMCountersReply, err error) {
	reply = new(GetDOMCountersReply)
	err = rpcc.Invoke(ctx, "Memory.getDOMCounters", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Memory", Op: "GetDOMCounters", Err: err}
	}
	return
}

// SetPressureNotificationsSuppressed invokes the Memory method. Enable/disable suppressing memory pressure notifications in all processes.
func (d *domainClient) SetPressureNotificationsSuppressed(ctx context.Context, args *SetPressureNotificationsSuppressedArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Memory.setPressureNotificationsSuppressed", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Memory.setPressureNotificationsSuppressed", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Memory", Op: "SetPressureNotificationsSuppressed", Err: err}
	}
	return
}

// SimulatePressureNotification invokes the Memory method. Simulate a memory pressure notification in all processes.
func (d *domainClient) SimulatePressureNotification(ctx context.Context, args *SimulatePressureNotificationArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Memory.simulatePressureNotification", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Memory.simulatePressureNotification", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Memory", Op: "SimulatePressureNotification", Err: err}
	}
	return
}