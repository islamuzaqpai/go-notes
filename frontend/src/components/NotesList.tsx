import { useEffect, useState } from "react";
import type { Note } from "../api";
import { getNotes, deleteNote, updateNote } from "../api";
import styles from "../styles/NotesList.module.css";

export default function NotesList() {
  const [notes, setNotes] = useState<Note[]>([]);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    loadNotes();
  }, []);

  const loadNotes = () => {
    getNotes()
      .then((data) => {
        if (Array.isArray(data)) {
          setNotes(data);
        } else {
          setNotes([]);
        }
      })
      .catch((err) => {
        console.error("Failed to load notes:", err);
        setError("Failed to load notes");
      });
  };

  const handleDelete = async (id: number) => {
    if (!confirm("Are you sure you want to delete this note?")) return;
    await deleteNote(id);
    loadNotes();
  };

  const handleUpdate = async (note: Note) => {
    const newTitle = prompt("Edit title:", note.title);
    const newContent = prompt("Edit content:", note.content);
    if (newTitle !== null && newContent !== null) {
      await updateNote(note.id, { title: newTitle, content: newContent });
      loadNotes();
    }
  };

  if (error) {
    return <p className="text-red-500">{error}</p>;
  }

  return (
    <div className={styles.container}>
      <h2 className={styles.title}>Notes</h2>
      {notes.length > 0 ? (
        notes.map((n) => (
          <div key={n.id} className={styles.noteCard}>
            <strong className={styles.noteTitle}>{n.title}</strong>
            <p className={styles.noteContent}>{n.content}</p>
            <div className={styles.actions}>
              <button
                className={`${styles.button} ${styles.buttonEdit}`}
                onClick={() => handleUpdate(n)}
              >
                Edit
              </button>
              <button
                className={`${styles.button} ${styles.buttonDelete}`}
                onClick={() => handleDelete(n.id)}
              >
                Delete
              </button>
            </div>
          </div>
        ))
      ) : (
        <p>No notes found.</p>
      )}
    </div>
  );
}
