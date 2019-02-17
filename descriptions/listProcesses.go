package descriptions

// ListProcesses :
const ListProcesses = `SSHes into a node and runs "systemctl | grep"

EXAMPLES
========

$ waggle list-processes -n 004
{
  "edge processor": [
    " wagglerw.mount loaded active mounted /wagglerw ",
    " rabbitmq-server.service loaded active running RabbitMQ broker ",
    " waggle-heartbeat.service loaded active running Triggers Wagman heartbeat line. ",
    " waggle-image-exporter.service loaded activating auto-restart Image Exporter ",
    " waggle-image-producer.service loaded activating auto-restart Image Producer ",
    " waggle-monitor-plugins.service loaded active running Monitors plugins. ",
    " waggle-plugin-50-0.0.0-0.service loaded active running Car and pedestraian counter plugin. ",
    " waggle-plugin-message-router.service loaded active running Routes messages to plugins within a subsystem. ",
    " waggle-core.target loaded active active Waggle Core ",
    " waggle-platform.target loaded active active Waggle Platform ",
    " waggle-monitor.timer loaded active waiting Automatically correct system state. "
  ],
  "node controller": [
    "wagglerw.mount loaded active mounted /wagglerw ",
    "rabbitmq-server.service loaded active running RabbitMQ broker ",
    "waggle-epoch.service loaded active running Maintains the date and time on the node. ",
    "waggle-heartbeat.service loaded active running Triggers Wagman heartbeat line. ",
    "waggle-monitor-connectivity.service loaded active running Monitors node controller connectivity status. ",
    "waggle-monitor-plugins.service loaded active running Monitors plugins. ",
    "waggle-monitor-shutdown.service loaded active running Monitors shutdown signals. ",
    "waggle-node-message-router.service loaded active running Routes messages to subsystems within a node. ",
    "waggle-plugin-coresense.service loaded active running Coresense 4.1 plugin. ",
    "waggle-plugin-message-router.service loaded active running Routes messages to plugins within a subsystem. ",
    "waggle-reverse-tunnel.service loaded active running Maintains an SSH reverse tunnel on Beehive. ",
    "waggle-stage-messages-ep.service loaded active running Stage messages from edge processor. ",
    "waggle-stage-messages-nc.service loaded active running Stage messages from node controller. ",
    "waggle-wagman-driver.service loaded active running Wagman Driver ",
    "waggle-wwan.service loaded active running ATT WWAN Client ",
    "waggle-core.target loaded active active Waggle Core ",
    "waggle-platform.target loaded active active Waggle Platform ",
    "waggle-status.timer loaded active waiting Waggle status every 5 minutes "
  ]
}
`
