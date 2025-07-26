import { useState } from "react";
import { addNote } from "../api";
import styles from "../styles/AddNoteForm.module.css";

interface Props {
  onAdd: () => void;
}

export default function AddNoteForm({ onAdd }: Props) {
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");

  const submit = async (e: React.FormEvent) => {
    e.preventDefault();
    await addNote({ title, content });
    setTitle("");
    setContent("");
    onAdd();
  };

  return (
    <form onSubmit={submit} className={styles.container}>
      <h2 className={styles.title}>Add New Note</h2>

      <div className={styles.formGroup}>
        <label className={styles.label}>Title</label>
        <input
          className={styles.input}
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          placeholder="Title"
        />
      </div>

      <div className={styles.formGroup}>
        <label className={styles.label}>Content</label>
        <textarea
          className={styles.textarea}
          value={content}
          onChange={(e) => setContent(e.target.value)}
          placeholder="Content"
        />
      </div>

      <button type="submit" className={styles.button}>
        Add Note
      </button>
    </form>
  );
}
