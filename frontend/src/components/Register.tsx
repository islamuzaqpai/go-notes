import { useState } from "react";
import axios from "axios";
import styles from "../styles/Register.module.css";

type RegisterProps = {
  onRegister: () => void;
};

const Register = ({ onRegister }: RegisterProps) => {
  const [username, setUsername] = useState(""); 
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleRegister = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      await axios.post("http://localhost:8080/register", {
        username, 
        email,
        password,
      });

      alert("Registration successful!");
      onRegister();
    } catch (err) {
      alert("Registration failed.");
      console.error(err);
    }
  };

  return (
    <form onSubmit={handleRegister} className={styles.form}>
      <h1 className={styles.title}>Register</h1>
      <input
        type="text"
        placeholder="Username"
        className={styles.input}
        value={username}
        onChange={(e) => setUsername(e.target.value)}
      />
      <input
        type="email"
        placeholder="Email"
        className={styles.input}
        value={email}
        onChange={(e) => setEmail(e.target.value)}
      />
      <input
        type="password"
        placeholder="Password"
        className={styles.input}
        value={password}
        onChange={(e) => setPassword(e.target.value)}
      />
      <button type="submit" className={styles.button}>
        Register
      </button>
    </form>
  );
};

export default Register;
