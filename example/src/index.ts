import express, { Request, Response } from "express";
import { Pool } from "pg";

const app = express();

app.use(express.json());

const pool = new Pool({
  host: process.env.DB_HOST,
  port: Number(process.env.DB_PORT),
  user: process.env.DB_USER,
  password: process.env.DB_PASSWORD,
  database: process.env.DB_NAME,
});

async function init() {
  await pool.query(`
    CREATE TABLE IF NOT EXISTS notes (
      id SERIAL PRIMARY KEY,
      title TEXT NOT NULL
    )
  `);
}

app.get("/notes", async (_: Request, res: Response) => {
  const result = await pool.query(
      "SELECT * FROM notes ORDER BY id DESC"
  );

  res.json(result.rows);
});

app.post("/notes", async (req: Request, res: Response) => {
  const { title } = req.body;

  const result = await pool.query(
      "INSERT INTO notes (title) VALUES ($1) RETURNING *",
      [title]
  );

  res.json(result.rows[0]);
});

init().then(() => {
  app.listen(process.env.PORT, () => {
    console.log(`Server running on port ${process.env.PORT}`);
  });
});