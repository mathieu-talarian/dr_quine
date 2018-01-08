#include <stdio.h>
#include <stdlib.h>
#include <string.h>
int main(int ac, char **av) {
int x = 0;
if (!strstr(av[0], "_")) x++;
ac = x;
if (x <= 0) return 0;char *file;char *bin;char *exe;
asprintf(&file, "Sully_%d.c", x-1);
asprintf(&bin, "clang -Wall -Werror -Wextra Sully_%1$d.c -o Sully_%1$d;", x-1);
asprintf(&exe, "./Sully_%d", x-1);
FILE *fp = fopen(file, "w");
char *c = "#include <stdio.h>%1$c#include <stdlib.h>%1$c#include <string.h>%1$cint main(int ac, char **av) {%1$cint x = %2$d;%1$cif (!strstr(av[0], %3$c_%3$c)) x++;%1$cac = x;%1$cif (x <= 0) return 0;char *file;char *bin;char *exe;%1$casprintf(&file, %3$cSully_%%d.c%3$c, x-1);%1$casprintf(&bin, %3$cclang -Wall -Werror -Wextra Sully_%%1$d.c -o Sully_%%1$d;%3$c, x-1);%1$casprintf(&exe, %3$c./Sully_%%d%3$c, x-1);%1$cFILE *fp = fopen(file, %3$cw%3$c);%1$cchar *c = %3$c%4$s%3$c;%1$cfprintf(fp, c, 10, x - 1, 34, c);fclose(fp);system(bin);system(exe);free(bin);free(exe);bin = NULL;exe = NULL;free(file);file = NULL;return 0;%1$c};%1$c";
fprintf(fp, c, 10, x - 1, 34, c);fclose(fp);system(bin);system(exe);free(bin);free(exe);bin = NULL;exe = NULL;free(file);file = NULL;return 0;
};
