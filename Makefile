# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
# the version is hold in the file .version
LXX_SOURCE_FOLDER=tlg0527
LXX_TARGET_FOLDER=First1KGreek-LXX-XML


.PHONY: test help
.DEFAULT_GOAL := help

download-source:	## download the source files from First 1k years of greek
	@echo "downloading source files from First1KGreek repo..."
	@wget https://github.com/OpenGreekAndLatin/First1KGreek/archive/master.zip -O src/First1KGreek.zip
	@echo "extract LXX and remove other works..."
	@unzip -qq src/First1KGreek.zip -d src
	@cp -r src/First1KGreek-master/data/${LXX_SOURCE_FOLDER} src/${LXX_TARGET_FOLDER}
	@echo "cleanup"
	@rm -rf src/First1KGreek-master
	@rm -rf src/First1KGreek-master.zip
	@echo "🍺 done"

remove-xml-tags:	## remove tags that hurt processing of texts
	@find src/First1KGreek-LXX -name "*.xml" | xargs sed -i '' 's/<l>//g'
	@find src/First1KGreek-LXX -name "*.xml" | xargs sed -i '' 's/<\/l>//g'
	@find src/First1KGreek-LXX -name "*.xml" | xargs sed -i '' 's/<lg>//g'
	@find src/First1KGreek-LXX -name "*.xml" | xargs sed -i '' 's/<\/lg>//g'
	@find src/First1KGreek-LXX -name "*.xml" | xargs sed -i '' 's/<l rend="indent">//g'

parse-xml: ## parse XML
	@echo "==> parsing source files and extracting text"
	@go run scripts/parsing/parse_lxx.go

help:
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
