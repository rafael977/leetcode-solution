.PHONY: arrange
arrange: goName = $(shell ls $(SN)*.go | head -1)
arrange:
	dirname=$(basename $(goName)); \
	mkdir $$dirname; \
	mv $${dirname}.go solution.go; \
	mv solution.go $${dirname}; \
	code $${dirname}/solution.go