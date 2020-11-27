# Create the list of directories
C_SOURCEDIRS = hello1/lib hello2/lib
GO_SOURCEDIRS = hello1 hello2

C_SOURCES = $(foreach dir,$(C_SOURCEDIRS),$(wildcard $(dir)/*.c))
GOPLUGINS_SOURCES = $(foreach dir,$(C_SOURCEDIRS),$(wildcard $(dir)/*.go))
GOPLUGINS_SHARED = $(foreach dir,$(GO_SOURCEDIRS),$(dir)/$(dir).so)
C_OBJS := $(C_SOURCES:.c=.o)
C_ARCHIVES := $(join $(dir $(C_SOURCES)), $(addprefix lib, $(notdir $(C_SOURCES:.c=.a))))
TARGET = main

CC = gcc
RM = rm -rf
AR = ar
GOBUILD = go build

%.o: %.c
	$(CC) -fvisibility=hidden -c $< -o $@

lib%.a: %.o
	$(AR) rcs $@ $<

%.so: $(C_ARCHIVES) $(GOPLUGINS_SOURCES)
	cd $(dir $@); $(GOBUILD) -buildmode=plugin -o $(notdir $@)

$(TARGET): $(GOPLUGINS_SHARED) main.go
	$(GOBUILD) -o $(TARGET)

.PHONY: all clean

all: $(C_OBJS) $(C_ARCHIVES) $(GOPLUGINS_SHARED) $(TARGET)

clean:
	$(RM) $(C_OBJS) $(C_ARCHIVES) $(GOPLUGINS_SHARED) $(TARGET)
