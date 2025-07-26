import axios from "axios";

const API_URL = "http://localhost:8080";

let token = "";

export const setToken = (t: string) => {
  token = t;
  localStorage.setItem("token", t);
};

const getToken = () => {
  if (!token) {
    token = localStorage.getItem("token") || "";
  }
  return token;
};

export const login = async (email: string, password: string) => {
  const res = await axios.post(`${API_URL}/login`, { email, password });
  return res.data; // { token: string }
};

export type Note = {
  id: number;
  title: string;
  content: string;
  created: string;
};

export const getNotes = async (): Promise<Note[]> => {
  try {
    const res = await axios.get(`${API_URL}/notes`, {
      headers: { Authorization: `Bearer ${getToken()}` },
    });
    return res.data;
  } catch (err) {
    console.error("Failed to fetch notes", err);
    return [];
  }
};

export const addNote = async (note: { title: string; content: string }) => {
  const res = await axios.post(`${API_URL}/notes`, note, {
    headers: { Authorization: `Bearer ${getToken()}` },
  });
  return res.data;
};

export const deleteNote = async (id: number) => {
  try {
    await axios.delete(`${API_URL}/notes/${id}`, {
      headers: { Authorization: `Bearer ${getToken()}` },
    });
  } catch (err) {
    console.error(`Failed to delete note with id ${id}`, err);
    throw err;
  }
};

export const updateNote = async (
  id: number,
  updatedNote: { title: string; content: string }
) => {
  try {
    await axios.put(`${API_URL}/notes/${id}`, updatedNote, {
      headers: { Authorization: `Bearer ${getToken()}` },
    });
  } catch (err) {
    console.error(`Failed to update note with id ${id}`, err);
    throw err;
  }
};
