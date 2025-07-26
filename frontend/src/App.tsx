import { useState } from "react";
import LoginForm from "./components/LoginForm";
import Register from "./components/Register";
import NotesList from "./components/NotesList";
import AddNoteForm from "./components/AddNoteForm";

function App() {
  const [loggedIn, setLoggedIn] = useState(false);
  const [showRegister, setShowRegister] = useState(false);
  const [update, setUpdate] = useState(false);

  if (!loggedIn) {
    return showRegister ? (
      <>
        <Register onRegister={() => setShowRegister(false)} />
        <p className="text-center mt-4">
          Already have an account?{" "}
          <button className="text-blue-500" onClick={() => setShowRegister(false)}>
            Log In
          </button>
        </p>
      </>
    ) : (
      <>
        <LoginForm onLogin={() => setLoggedIn(true)} />
        <p className="text-center mt-4">
          Don't have an account?{" "}
          <button className="text-blue-500" onClick={() => setShowRegister(true)}>
            Register
          </button>
        </p>
      </>
    );
  }

  return (
    <>
      <AddNoteForm onAdd={() => setUpdate(!update)} />
      <NotesList key={update ? "1" : "0"} />
    </>
  );
}

export default App;
