var MainBox  = React.createClass({
    render:function(){
        return(
            <App/>
        );
    }
});
var App = React.createClass({
    //setting up initial state
    getInitialState:function(){
        return{
            data:[]
        };
    },
    componentDidMount(){
        this.getDataFromServer('http://localhost:8080/api/v1/sma200');
    },
    //showResult Method
        showResult: function(response) {

            this.setState({
                data: response
            });
    },
    //making ajax call to get data from server
    getDataFromServer:function(URL){
        $.ajax({
            type:"GET",
            dataType:"json",
            url:URL,
            success: function(response) {
                this.showResult(response);
            }.bind(this),
            error: function(xhr, status, err) {
                console.error(this.props.url, status, err.toString());
            }.bind(this)
        });
    },
    render:function(){
        return(
            <div>
                <Result result={this.state.data}/>
            </div>
        );
    }
});

var Result = React.createClass({
    render:function(){
        var result = this.props.result.map(function(result,index){
            return <ResultItem key={index} user={ result } />
            });
        return(
            <div className="row">
                <h4 align="left">Stocks Breaking 200 day simple moving average</h4>
                <br></br>
                <div className="col-md-12">
                    <table className="table table-bordered">
                        <thead>
                            <tr>
                                <th className="col-md-4">Stock</th>
                                <th >Status</th>
                                <th>Link</th>
                            </tr>
                        </thead>
                        <tbody>
                            {result}
                        </tbody>
                    </table>
                </div>
            </div>
        );
    }
});
var ResultItem = React.createClass({
    render:function(){
        var fin = this.props.user;
        return(
            <tr >
                <td>{fin.Stock}</td>
                <td><img src={fin.Status} style={{width: '15px', height: '15px'}}></img></td>
                <td><a href={fin.Url}>{fin.Stock}</a></td>
            </tr>
        );
    }
});

ReactDOM.render(
    <MainBox />,
    document.querySelector("#app")
);