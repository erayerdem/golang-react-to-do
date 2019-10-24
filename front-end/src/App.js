import React, {Component} from 'react';
import "./App.css"
import axios from 'axios'


class App extends Component {


    constructor(props) {

        super(props);
        this.state = {
            host: "http://localhost:8080",
            doing: "",
            doinglist: []
        }
    }

    addtoTodo = () => {

        axios.post(
            this.state.host + "/todo", {
                todo: this.state.doing,
            })
            .then((response) => {

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
                        <button value={index} onClick={(event => {

                            this.setState({
                                    doinglist: this.state.doinglist.filter((value1, index1) =>
                                        !(index1 == event.target.value)
                                    )
                                }
                            )
                            console.log(this.state.doinglist)
                            event.preventDefault();
                        })}>Done
                        </button>
                        {value}</li>
                ))}
            </div>
        );
    }

}

export default App;