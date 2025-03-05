[GO__BADGE]: https://img.shields.io/badge/goLang-fff?style=for-the-badge&logo=go
[AWS_BADGE]: https://img.shields.io/badge/Fly.io-aa9bd2.svg?style=for-the-badge&logo=fly.io&logoColor=white

<h1 align="center" style="font-weight: bold;">Priagram Backend</h1>

<div align="center">

![golang][GO__BADGE]
![AWS][AWS_BADGE]

</div>

<p align="center">
 <a href="#started">Getting Started</a> • 
  <a href="#routes">API Endpoints</a> •
 <a href="#colab">Collaborators</a> •
 <a href="#contribute">Contribute</a>
</p>

<p align="center">
  <b>An API to read Prisma models and transform it into a diagram.</b>
</p>

<h2 id="started">🚀 Getting started</h2>

<h3>Prerequisites</h3>

You will need Golang with version 1.23.4 to run this project

- [Git](https://git.com)
- [Golang](https://go.dev/doc/install)

<h3>Cloning</h3>

How to clone the project:

- first you need to fork it clicking in the fork button on github
- then you need to clone your fork

```bash
git clone https://github.com/<your-github-username>/priagram-backend.git
```

<h3>Starting</h3>

How to start your project

```bash
cd priagram-backend
go run src/main.go
```

<h2 id="routes">📍 API Endpoints</h2>

Here you can list the main routes of your API, and what are their expected request bodies.
​
| route | description  
|----------------------|-----------------------------------------------------
| <kbd>GET /api/healthcheck</kbd> | returns 200 with a message checking if he api is ok [response details](#get-healthcheck)
| <kbd>POST /api/prisma</kbd> | receive prisma models table and returns nodes with the model name, columns, enums and the relations [request details](#post-prisma-diagram)

<h3 id="get-healthcheck">GET /api/healthcheck</h3>

**RESPONSE**

```json
{
  "status": 200,
  "message": "application working correctly"
}
```

<h3 id="post-prisma-diagram">POST /api/prisma</h3>

**REQUEST**

```json
{
  "source": " enum TestEnum {\n  TESTNUM\n  ANOTHERENUM\n}\n\nenum UserRole {\n  SUPER\n  ADMIN\n  AGENT\n  BI_ANALYST\n}\n\n model User {\n  id            String    @id @default(cuid())\n  name          String?\n  email         String?   @unique\n  emailVerified DateTime?\n  password      String?   @default(\"\")\n  image         String?\n  phone         String?\n   test          TestEnum\n  role          UserRole\n activeStatus  Boolean   \n\n  createdAt DateTime  @default(now())\n  updatedAt DateTime  @updatedAt\n  deletedAt DateTime?\n\n  tickets  TicketsToAgents[]\n  Messages Messages[]\n}\n"
}
```

**RESPONSE**

```json
{
  "data": [
    {
      "id": "testenum",
      "type": "prisma-table",
      "data": {
        "modelName": "testenum",
        "modelContent": [
          {
            "id": "TESTNUM->enum",
            "name": "TESTNUM",
            "type": "enum"
          },
          {
            "id": "ANOTHERENUM->enum",
            "name": "ANOTHERENUM",
            "type": "enum"
          }
        ]
      }
    },
    {
      "id": "userrole",
      "type": "prisma-table",
      "data": {
        "modelName": "userrole",
        "modelContent": [
          {
            "id": "SUPER->enum",
            "name": "SUPER",
            "type": "enum"
          },
          {
            "id": "ADMIN->enum",
            "name": "ADMIN",
            "type": "enum"
          },
          {
            "id": "AGENT->enum",
            "name": "AGENT",
            "type": "enum"
          },
          {
            "id": "BI_ANALYST->enum",
            "name": "BI_ANALYST",
            "type": "enum"
          }
        ]
      }
    },
    {
      "id": "user",
      "type": "prisma-table",
      "data": {
        "modelName": "user",
        "modelContent": [
          {
            "id": "id->String",
            "name": "id",
            "type": "String"
          },
          {
            "id": "name->String?",
            "name": "name",
            "type": "String?"
          },
          {
            "id": "email->String?",
            "name": "email",
            "type": "String?"
          },
          {
            "id": "emailVerified->DateTime?",
            "name": "emailVerified",
            "type": "DateTime?"
          },
          {
            "id": "password->String?",
            "name": "password",
            "type": "String?"
          },
          {
            "id": "image->String?",
            "name": "image",
            "type": "String?"
          },
          {
            "id": "phone->String?",
            "name": "phone",
            "type": "String?"
          },
          {
            "id": "test->TestEnum",
            "name": "test",
            "type": "TestEnum"
          },
          {
            "id": "role->UserRole",
            "name": "role",
            "type": "UserRole"
          },
          {
            "id": "activeStatus->Boolean",
            "name": "activeStatus",
            "type": "Boolean"
          },
          {
            "id": "createdAt->DateTime",
            "name": "createdAt",
            "type": "DateTime"
          },
          {
            "id": "updatedAt->DateTime",
            "name": "updatedAt",
            "type": "DateTime"
          },
          {
            "id": "deletedAt->DateTime?",
            "name": "deletedAt",
            "type": "DateTime?"
          },
          {
            "id": "tickets->TicketsToAgents[]",
            "name": "tickets",
            "type": "TicketsToAgents[]"
          }
        ]
      }
    }
  ],
  "relations": [
    {
      "id": "user->TestEnum",
      "source": "testenum",
      "target": "user",
      "sourceHandle": "testenum-source",
      "targetHandle": "user-target",
      "animated": true
    },
    {
      "id": "user->UserRole",
      "source": "userrole",
      "target": "user",
      "sourceHandle": "userrole-source",
      "targetHandle": "user-target",
      "animated": true
    }
  ]
}
```

<h2 id="colab">🤝 Collaborators</h2>

<table>
  <tr>
    <td align="center">
      <a href="#">
        <img src="https://avatars.githubusercontent.com/u/103784814?v=4" width="100px;" alt="victor hugo Profile Picture"/><br>
        <sub>
          <b>Victor Hugo - hxsggsz</b>
        </sub>
      </a>
    </td>
  </tr>
</table>
