--liquibase formatted sql

--changeset Piyawat:1 labels:go-service-client context:centralize-configuration
--comment: Insert SERVICE_SERVER_URL
INSERT INTO centralize_configuration.properties (application, profile, `key`, value) VALUES
('go-service-client', 'default', 'host.go_service_server.url', 'http://localhost:8111/go-service-server');
INSERT INTO centralize_configuration.properties (application, profile, `key`, value) VALUES
('go-service-client', 'default', 'gin_mode', 'debug');
INSERT INTO centralize_configuration.properties (application, profile, `key`, value) VALUES
('go-service-client', 'default', 'server_address', ':8112');
--rollback DELETE FROM centralize_configuration.properties WHERE application='go-service-client' AND `key`='host.go_service_server.url';
--rollback DELETE FROM centralize_configuration.properties WHERE application='go-service-client' AND `key`='gin_mode';
--rollback DELETE FROM centralize_configuration.properties WHERE application='go-service-client' AND `key`='server_address';
