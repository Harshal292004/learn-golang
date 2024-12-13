import Header from "./components/Header";
import Display from "./components/Display";
import Input from "./components/Input";
const App = () => {
  
  return (
    <div className="min-h-screen bg-zinc-800 text-white flex flex-col">
      <Header />
      <Display />
      <Input />
    </div>
  );
};

export default App;
