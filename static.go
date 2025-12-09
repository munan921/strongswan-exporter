package main

import (
	"fmt"
	"html"
)

func getStaticHTML(config *Config) (string, error) {
	sessions, err := getSessions(config)
	if err != nil {
		return "", err
	}

	htmlContent := `<!DOCTYPE html>
<html>
<head>
	<title>StrongSwan Active Sessions</title>
	<style>
		body {
			font-family: Arial, sans-serif;
			margin: 20px;
			background-color: #f5f5f5;
		}
		h1 {
			color: #333;
		}
		table {
			width: 100%;
			border-collapse: collapse;
			background-color: white;
			box-shadow: 0 2px 4px rgba(0,0,0,0.1);
		}
		th {
			background-color: #4CAF50;
			color: white;
			padding: 12px;
			text-align: left;
			font-weight: bold;
		}
		td {
			padding: 10px;
			border-bottom: 1px solid #ddd;
		}
		tr:hover {
			background-color: #f5f5f5;
		}
		.no-sessions {
			padding: 20px;
			text-align: center;
			color: #666;
		}
		.stats {
			margin: 20px 0;
			padding: 15px;
			background-color: white;
			border-radius: 5px;
			box-shadow: 0 2px 4px rgba(0,0,0,0.1);
		}
	</style>
</head>
<body>
	<h1>StrongSwan Active Sessions</h1>
	<div class="stats">
		<strong>Total Sessions:</strong> ` + fmt.Sprintf("%d", len(sessions)) + `
	</div>
`

	if len(sessions) == 0 {
		htmlContent += `<div class="no-sessions">No active sessions</div>`
	} else {
		htmlContent += `<table>
		<thead>
			<tr>
				<th>Server</th>
				<th>Protocol</th>
				<th>State</th>
				<th>Remote Host</th>
				<th>Remote Port</th>
				<th>Remote ID</th>
				<th>Remote TS</th>
				<th>Local TS</th>
				<th>Established</th>
				<th>Bytes In</th>
				<th>Bytes Out</th>
			</tr>
		</thead>
		<tbody>
`

		for _, session := range sessions {
			htmlContent += fmt.Sprintf(`			<tr>
				<td>%s</td>
				<td>%s</td>
				<td>%s</td>
				<td>%s</td>
				<td>%s</td>
				<td>%s</td>
				<td>%s</td>
				<td>%s</td>
				<td>%s</td>
				<td>%s</td>
				<td>%s</td>
			</tr>
`,
				html.EscapeString(session.Server),
				html.EscapeString(session.Protocol),
				html.EscapeString(session.State),
				html.EscapeString(session.RemoteHost),
				html.EscapeString(session.RemotePort),
				html.EscapeString(session.RemoteID),
				html.EscapeString(session.RemoteTS),
				html.EscapeString(session.LocalTS),
				html.EscapeString(session.Established),
				html.EscapeString(session.BytesIn),
				html.EscapeString(session.BytesOut),
			)
		}

		htmlContent += `		</tbody>
	</table>
`
	}

	htmlContent += `</body>
</html>`

	return htmlContent, nil
}