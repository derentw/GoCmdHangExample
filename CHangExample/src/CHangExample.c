/*
 ============================================================================
 Name        : CHangExample.c
 Author      : 
 Version     :
 Copyright   : Your copyright notice
 Description : Hello World in C, Ansi-style
 ============================================================================
 */

#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>

#define BUF_SIZE 1024

void* thread_copy() {
	FILE* fp_src = NULL;
	FILE* fp_dest = NULL;
	char buf[BUF_SIZE];
	size_t numRead = 0;
	fp_src = fopen("srcFile", "r");
	if (fp_src == NULL) {
		printf("fopen(srcFile, r) fail\n");
		return NULL;
	}

	fp_dest = fopen("/mnt/fuse/destFile", "w");
	if (fp_dest == NULL) {
		printf("fopen(destFile, w) fail\n");
		return NULL;
	}
	size_t offset = 0;
	while ((numRead = fread(buf, 1, BUF_SIZE, fp_src)) > 0) {
		printf("write offset=%ld start\n", offset);
		if (fwrite(buf, 1, BUF_SIZE, fp_dest) != numRead) {
			printf("fwrite() fail\n");
			break;
		}
		printf("write offset=%ld done\n", offset);
		offset += numRead;
	}
	fclose(fp_src);
	fclose(fp_dest);
	return NULL;
}
void* thread_system() {
	int i = 0;
	for (i = 0; i < 10; i++) {
		printf("systemctl is-active smb start loop=%d\n", i);
		system("systemctl is-active smb");
		printf("systemctl is-active smb done  loop=%d\n", i);
	}
	return NULL;
}
int main(void) {
	system("dd if=/dev/zero of=srcFile bs=1MiB count=1");
	pthread_t pid_copy;
	pthread_create(&pid_copy, NULL, thread_copy, NULL);

	pthread_t pid_system;
	pthread_create(&pid_system, NULL, thread_system, NULL);

	pthread_join(pid_copy, NULL);
	pthread_join(pid_system, NULL);
	return EXIT_SUCCESS;
}
