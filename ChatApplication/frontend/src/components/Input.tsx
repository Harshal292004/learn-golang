import React, { useContext, useState } from "react";
import { AppContext } from "../context/AppContext";

const Input = () => {
  const [message, setMessage] = useState("");
  const context = useContext(AppContext);

  if (!context) {
    return (
      <div className="min-h-screen bg-zinc-800 text-red-200 text-center flex justify-center items-center font-bold text-lg">
        There was an error loading the chat context.
      </div>
    );
  }

  const { sendMessage } = context;

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (message.trim()) {
      sendMessage(message);
      setMessage("");
    }
  };

  return (
    <div className="p-4">
      <form onSubmit={handleSubmit} className="flex">
        <input
          className="w-full rounded-lg bg-slate-600 text-white p-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
          type="text"
          onChange={(e) => setMessage(e.target.value)}
          value={message}
          placeholder="Type a message..."
        />
        <button
          type="submit"
          className="ml-2 p-2 bg-blue-600 text-white rounded-lg hover:bg-blue-500 transition duration-300"
        >
          Send
        </button>
      </form>
    </div>
  );
};

export default Input;
