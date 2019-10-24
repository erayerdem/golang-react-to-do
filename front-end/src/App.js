import React, {Component} from 'react';
import "./App.css"
import axios from 'axios'


class App extends Component {

    componentDidMount() {
        axios.get(
            this.state.host + "/todo")
            .then((response) => {
                console.log(response)
                if (response.status === 200 && response.data.Todo) {
                    this.setState({
                        doinglist: response.data.Todo
                    })
                }
            })
            .catch(function (error) {
                console.log(error);
            });
    }

    constructor(props) {

        super(props);
        this.state = {
            host: "http://localhost:8080",
            doing: "",
            doinglist: []
        }
    }

    addtoTodo = () => {
        console.info(this.state.doinglist);
        axios.post(
            this.state.host + "/todo", {
                todo: this.state.doing,
            })
            .then((response) => {
                console.log(response)
                if (response.status === 201) {


                    this.setState({
                        doinglist: [...this.state.doinglist, this.state.doing]

                    })
                }


            })
            .catch(function (error) {
                console.log(error);
            });
    };

    done = (event) => {


        let value = parseInt(event.target.value);
        console.log(this.state.doinglist[value]);
        axios.post(
            this.state.host + "/deletetodo", {
                todo: this.state.doinglist[value],
            })
            .then((response) => {

                if (response.status === 200) {
                    console.log("delete location")
                    this.setState({
                            doinglist: this.state.doinglist.filter((value1, index1) =>
                                !(index1 === value)
                            )
                        }
                    )
                }
            })
            .catch(function (error) {
                console.log(error);
            });
    };

    render() {
        return (
            <div className={"divir"}>
                <input id={"todo"} type={"text"} value={this.state.doing}
                       onChange={(event) => this.setState({doing: event.target.value})}
                       onKeyDown={event => {
                           if (event.keyCode === 13) {
                               this.addtoTodo();
                           }
                       }}>
                </input>

                {this.state.doinglist.map((value, index) => (

                    <li key={index}>
                        <button value={index} onClick={this.done}>Done
                        </button>
                        {value}</li>
                ))}

            </div>
        );
    }

}

export default App;