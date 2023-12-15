.PHONY: liaz-admin liaz-business liaz-oauth compile clean

RM			:= rm -rf

PROJECT_DIR	:= $(shell pwd)
BUILD		:= build

compile:
	@echo "cd $(PROJECT_DIR)/admin && CGO_ENABLED=0 go build -ldflags="-w -s" -a -installsuffix "" -gcflags="all=-N -l" -o $(PROJECT_DIR)/$(BUILD)/liaz-admin $(PROJECT_DIR)/admin/main.go" 
	@cd $(PROJECT_DIR)/admin && CGO_ENABLED=0 go build -ldflags="-w -s" -a -installsuffix "" -gcflags="all=-N -l" -o $(PROJECT_DIR)/$(BUILD)/liaz-admin $(PROJECT_DIR)/admin/main.go	
	@echo "cd $(PROJECT_DIR)/business && CGO_ENABLED=0 go build -ldflags="-w -s" -a -installsuffix "" -gcflags="all=-N -l" -o $(PROJECT_DIR)/$(BUILD)/liaz-business $(PROJECT_DIR)/business/main.go" 
	@cd $(PROJECT_DIR)/business && CGO_ENABLED=0 go build -ldflags="-w -s" -a -installsuffix "" -gcflags="all=-N -l" -o $(PROJECT_DIR)/$(BUILD)/liaz-business $(PROJECT_DIR)/business/main.go
	@echo "cd $(PROJECT_DIR)/oauth && CGO_ENABLED=0 go build -ldflags="-w -s" -a -installsuffix "" -gcflags="all=-N -l" -o $(PROJECT_DIR)/$(BUILD)/liaz-oauth $(PROJECT_DIR)/oauth/main.go" 
	@cd $(PROJECT_DIR)/oauth && CGO_ENABLED=0 go build -ldflags="-w -s" -a -installsuffix "" -gcflags="all=-N -l" -o $(PROJECT_DIR)/$(BUILD)/liaz-oauth $(PROJECT_DIR)/oauth/main.go

liaz-admin:
	@echo "cd $(PROJECT_DIR)/admin && CGO_ENABLED=0 go build -ldflags="-w -s" -a -installsuffix "" -gcflags="all=-N -l" -o $(PROJECT_DIR)/$(BUILD)/$@ $(PROJECT_DIR)/admin/main.go" 
	@cd $(PROJECT_DIR)/admin && CGO_ENABLED=0 go build -ldflags="-w -s" -a -installsuffix "" -gcflags="all=-N -l" -o $(PROJECT_DIR)/$(BUILD)/$@ $(PROJECT_DIR)/admin/main.go
	@echo Executing $@ complete!
	$(PROJECT_DIR)/$(BUILD)/$@ start -e prod

liaz-business:
	@echo "cd $(PROJECT_DIR)/business && CGO_ENABLED=0 go build -ldflags="-w -s" -a -installsuffix "" -gcflags="all=-N -l" -o $(PROJECT_DIR)/$(BUILD)/$@ $(PROJECT_DIR)/business/main.go" 
	@cd $(PROJECT_DIR)/business && CGO_ENABLED=0 go build -ldflags="-w -s" -a -installsuffix "" -gcflags="all=-N -l" -o $(PROJECT_DIR)/$(BUILD)/$@ $(PROJECT_DIR)/business/main.go
	@echo Executing $@ complete!
	$(PROJECT_DIR)/$(BUILD)/$@ start -e prod

liaz-oauth:
	@echo "cd $(PROJECT_DIR)/oauth && CGO_ENABLED=0 go build -ldflags="-w -s" -a -installsuffix "" -gcflags="all=-N -l" -o $(PROJECT_DIR)/$(BUILD)/$@ $(PROJECT_DIR)/oauth/main.go" 
	@cd $(PROJECT_DIR)/oauth && CGO_ENABLED=0 go build -ldflags="-w -s" -a -installsuffix "" -gcflags="all=-N -l" -o $(PROJECT_DIR)/$(BUILD)/$@ $(PROJECT_DIR)/oauth/main.go
	@echo Executing $@ complete!
	$(PROJECT_DIR)/$(BUILD)/$@ start -e prod

clean:
	$(RM) $(PROJECT_DIR)/$(BUILD)
	$(RM) $(PROJECT_DIR)/cache
	$(RM) $(PROJECT_DIR)/log
	@echo Cleanup complete!