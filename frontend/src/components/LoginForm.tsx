import { useState } from "react";
import { login, setToken } from "../api";
import styles from "../styles/LoginForm.module.css";

interface Props {
  onLogin: () => void;
}

export default function LoginForm({ onLogin }: Props) {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const submit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const res = await login(email, password);
      setToken(res.token);
      onLogin();
    } catch {
      alert("Login failed");
    }
  };

  return (
    <form onSubmit={submit} className={styles.container}>
      <h2 className={styles.title}>Login</h2>

      <div className={styles.formGroup}>
        <label className={styles.label}>Email</label>
        <input
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          placeholder="Email"
          className={styles.input}
          type="email"
        />
      </div>

      <div className={styles.formGroup}>
        <label className={styles.label}>Password</label>
        <input
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          placeholder="Password"
          type="password"
          className={styles.input}
        />
      </div>

      <button type="submit" className={styles.button}>
        Login
      </button>
    </form>
  );
}
