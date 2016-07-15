#Montando a minha base de teste

```
#!/bin/bash
docker run --name some-cassandra-2 -v ./datadir:/var/lib/cassandra -d -p 0.0.0.0:7000:7000 -p 0.0.0.0:7001:7001 -p 0.0.0.0:9160:9160 -p 0.0.0.0:9042:9042 -p 0.0.0.0:7199:7199 cassandra
```

- Entrando na imagem teste
docker exec -it some-cassandra-2 bash

- Eexecutar dentro da imagem: cqlsh

- Criando uma keyspace
```
CREATE KEYSPACE dev WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'}  AND durable_writes = true;
```

Verificando como foi criada a keyspace
```
DESC KEYSPACE
```

```
use dev;
create table emp (empid int primary key, emp_first varchar, emp_last varchar, emp_dept varchar);

insert into emp (empid, emp_first, emp_last, emp_dept) values (1,'fred','smith','eng');
select * from emp;
update emp set emp_dept = 'fin' where empid = 1;
```

- "In Cassandra, if you want to query columns other than the primary key, you need to create a secondary index on them:"
```
create index idx_dept on emp(emp_dept);
select * from emp where emp_dept = 'fin';
```
