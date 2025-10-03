INSERT INTO sys_plugin (name, status, description)
VALUES ('logs', 'installed', '系统日志插件')
ON DUPLICATE KEY UPDATE status = 'installed';

INSERT INTO sys_menu (title, path, component, parent_id)
VALUES ('日志管理', '/sys/logs', 'views/logs/index.vue', 'sys')
ON DUPLICATE KEY UPDATE title = VALUES(title), component = VALUES(component);

INSERT INTO sys_api (path, method, description)
VALUES ('/logs/list', 'POST', '查询系统日志')
ON DUPLICATE KEY UPDATE description = VALUES(description);
