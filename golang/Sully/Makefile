NAME=Sully
GOCMD=go
GOBUILD=$(GOCMD) build
FLAGS=
SRCS=Sully.go


all: 
	@$(GOBUILD) -o $(NAME) $(SRCS)

$(NAME): 
	all

clean:
	@rm -rf $(NAME)

fclean: clean

re: fclean all

run: re
	@./$(NAME)

.PHONY: all
