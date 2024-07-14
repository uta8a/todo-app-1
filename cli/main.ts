const URL = Deno.env.get("URL") || "http://localhost:8000";
const TOKEN = Deno.env.get("GCLOUD_TOKEN");

type Raw = {
  done: boolean;
  content: string;
}

type Todo = {
  id: string;
  done: boolean;
  content: string;
}

const get = async () => {
  const res = await fetch(`${URL}`, {
    method: "GET",
    headers: {
      Authorization: `Bearer ${TOKEN}`,
    },
  });
  return res.json();
}

// await post({ done: false, content: "Buy milk" });
const post = async (raw: Raw) => {
  const res = await fetch(`${URL}`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${TOKEN}`,
    },
    body: JSON.stringify(raw),
  });
  return res.text();
}

// await put({ id: id, done: true, content: "Buy milk" });
const put = async (todo: Todo) => {
  const res = await fetch(`${URL}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${TOKEN}`,
    },
    body: JSON.stringify(todo),
  });
  return res.text();
}

// await del(id)
const del = async (id: string) => {
  const res = await fetch(`${URL}?id=${id}`, {
    method: "DELETE",
    headers: {
      Authorization: `Bearer ${TOKEN}`,
    },
  });
  return res.text();
}

const show = async () => {
  const data = await get();
  data.forEach((todo: Todo) => {
    console.log(`${todo.done ? "✅" : "⬜"} ${todo.content} (id: ${todo.id})`);
  })
}

Deno.args.forEach(async (arg) => {
  if (arg === "show") {
    await show();
  } else if (arg === "post") {
    await post({ done: false, content: Deno.args[2] });
  } else if (arg === "put") {
    await put({ id: Deno.args[2], done: Deno.args[3] === "true", content: Deno.args[4] });
  } else if (arg === "del") {
    await del(Deno.args[2]);
  } else if (arg === "done") {
    const id = Deno.args[2];
    const todo = (await get()).find((todo: Todo) => todo.id === id);
    if (todo) {
      await put({ id: id, done: true, content: todo.content });
    }
  }
})
