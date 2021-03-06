package sqlplugin

import (
	"database/sql"
	"time"
)

type (
	// VisibilityRow represents a row in executions_visibility table
	VisibilityRow struct {
		NamespaceID      string
		RunID            string
		WorkflowTypeName string
		WorkflowID       string
		StartTime        time.Time
		ExecutionTime    time.Time
		Status           int32
		CloseTime        *time.Time
		HistoryLength    *int64
		Memo             []byte
		Encoding         string
	}

	// VisibilitySelectFilter contains the column names within executions_visibility table that
	// can be used to filter results through a WHERE clause
	VisibilitySelectFilter struct {
		NamespaceID      string
		RunID            *string
		WorkflowID       *string
		WorkflowTypeName *string
		Status           int32
		MinTime          *time.Time
		MaxTime          *time.Time
		PageSize         *int
	}

	VisibilityDeleteFilter struct {
		NamespaceID string
		RunID       string
	}

	Visibility interface {
		// InsertIntoVisibility inserts a row into visibility table. If a row already exist,
		// no changes will be made by this API
		InsertIntoVisibility(row *VisibilityRow) (sql.Result, error)
		// ReplaceIntoVisibility deletes old row (if it exist) and inserts new row into visibility table
		ReplaceIntoVisibility(row *VisibilityRow) (sql.Result, error)
		// SelectFromVisibility returns one or more rows from visibility table
		// Required filter params:
		// - getClosedWorkflowExecution - retrieves single row - {namespaceID, runID, closed=true}
		// - All other queries retrieve multiple rows (range):
		//   - MUST specify following required params:
		//     - namespaceID, minStartTime, maxStartTime, runID and pageSize where some or all of these may come from previous page token
		//   - OPTIONALLY specify one of following params
		//     - workflowID, workflowTypeName, status (along with closed=true)
		SelectFromVisibility(filter VisibilitySelectFilter) ([]VisibilityRow, error)
		DeleteFromVisibility(filter VisibilityDeleteFilter) (sql.Result, error)
	}
)
