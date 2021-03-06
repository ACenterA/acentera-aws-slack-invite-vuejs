#
# == usage ==
#
#   - Launch local development
#     make dev
#

help:
	@printf '=======================================\n'
	@printf '[>] To launch local development \n'
	@printf '[CMD] make dev\n'
	@printf '=======================================\n'

dev: 
	@printf '[>] Provisioning the docker containers\n'                                        
	docker-compose -f docker-compose.yml up -d --build --force-recreate

down: 
	@printf '[>] Provisioning the docker containers\n'                                        
	docker-compose -f docker-compose.yml down
