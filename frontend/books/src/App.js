import './App.css';
import Information from "./Info.js";


function App() {
  return (
    <div className="App">
        <Information />
        <AddBook />
    </div>
  );
}




function AddBook() {

    const value = 'Book Title';
    return(
        <form>
            <label for="text-form"> Name of The Book: </label>
            <input type="text" value={value} id="text-form" />
        </form>
    )

}



export default App;
