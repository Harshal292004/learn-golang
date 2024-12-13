import { createContext, useEffect, useState } from "react";
interface Message {
    type: number;
    body: string;
}

interface AppContextType {
    messages: Message[];
    setMessages: (messages: Message[]) => void;
    sendMessage: (message: string) => void;
}

export const AppContext = createContext<AppContextType | null>(null);

const AppContextProvider = ({ children }: { children: React.ReactNode }) => {
    const [messages, setMessages] = useState<Message[]>([]); // List of chat messages
    const [ws, setWs] = useState<WebSocket | null>(null); // WebSocket connection state

    useEffect(() => {
        let websocket: WebSocket | null = null;
        let reconnectTimeout: NodeJS.Timeout;
    
        const connect = () => {
            try {
                websocket = new WebSocket("ws://localhost:9000/ws");
                
                websocket.onopen = () => {
                    console.log("WebSocket connection established successfully");
                    // Clear any existing reconnect timeout
                    if (reconnectTimeout) clearTimeout(reconnectTimeout);
                };
    
                websocket.onmessage = (event: MessageEvent) => {
                    try {
                        const message: Message = JSON.parse(event.data);
                        setMessages((prevMessages) => [...prevMessages, message]);
                    } catch (parseError) {
                        console.error("Error parsing message:", parseError);
                        console.error("Received data:", event.data);
                    }
                };
    
                websocket.onerror = (error) => {
                    console.error("WebSocket Error Details:", {
                        error,
                        readyState: websocket?.readyState,
                        url: websocket?.url
                    });
                };
    
                websocket.onclose = (event) => {
                    console.log(`WebSocket closed: 
                        Code: ${event.code}, 
                        Reason: ${event.reason}, 
                        Clean: ${event.wasClean}`);
                    
                    // Attempt to reconnect
                    reconnectTimeout = setTimeout(() => {
                        console.log("Attempting to reconnect...");
                        connect();
                    }, 3000); // 3 seconds delay
                };
    
                setWs(websocket);
            } catch (error) {
                console.error("WebSocket connection setup error:", error);
            }
        };
    
        connect();
    
        return () => {
            if (websocket) websocket.close();
            if (reconnectTimeout) clearTimeout(reconnectTimeout);
        };
    }, []);
    
    
    const sendMessage = (message: string) => {
        if (ws && ws.readyState === WebSocket.OPEN) {
            const messageObject: Message = {
                type: 1, // Just a placeholder, modify based on your message types
                body: message,
            };
            ws.send(JSON.stringify(messageObject)); // Send the message to the server
        } else {
            console.error("WebSocket is not connected.");
        }
    };

    const value = {
        messages,
        setMessages,
        sendMessage, 
    };

    return <AppContext.Provider value={value}>{children}</AppContext.Provider>;
};

export default AppContextProvider;
