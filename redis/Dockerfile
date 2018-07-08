FROM redis
COPY redis.conf /usr/local/etc/redis/redis.conf
RUN sysctl vm.overcommit_memory=1
RUN echo never > /sys/kernel/mm/transparent_hugepage/enabled
CMD [ "redis-server", "/usr/local/etc/redis/redis.conf" "--appendonly" "yes"]