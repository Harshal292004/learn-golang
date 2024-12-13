import { useContext } from "react";
import { AppContext } from "../context/AppContext";

const Display = () => {
  const context = useContext(AppContext);

  if (!context) {
    return (
      <div className="min-h-screen bg-zinc-800 text-red-200 text-center flex justify-center items-center font-bold text-lg">
        There was an error loading the messages.
      </div>
    );
  }

  const { messages } = context;

  return (
    <div className="flex-1 p-4 overflow-y-auto">
      {messages.length === 0 ? (
        <div className="text-center text-gray-400">No messages yet.</div>
      ) : (
        messages.map((message, index) => (
          <div
            key={index}
            className="rounded-lg p-4 text-sm bg-blue-300 text-black font-light mb-2"
          >
            {message.body}
          </div>
        ))
      )}
    </div>
  );
};

export default Display;
