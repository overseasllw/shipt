import React, { Component } from 'react';
import {Container,Menu,Icon,Tab,Table,Grid,Input,Button,Form} from 'semantic-ui-react';
import * as moment from 'moment';

class Display extends Component{

  constructor(props){
    super(props)
    this.state = {
      CategorySale:[],
      CustomerId :0,
      CustomerOrders:[],
      Type:"day",
      StartDate:moment().subtract(1, "days").format("YYYY-MM-DD"),
      EndDate:moment().format("YYYY-MM-DD"),
      ProductSales:[],
    }
    this.loadCustomerOrder = this.loadCustomerOrder.bind(this)
    this.setCustomerId = this.setCustomerId.bind(this)
    this.changeType = this.changeType.bind(this)
    this.setEndDate = this.setEndDate.bind(this)
    this.setStartDate = this.setStartDate.bind(this)
    this.loadProductSale = this.loadProductSale.bind(this)
    this.exportProductSale = this.exportProductSale.bind(this)
  }

  loadProductSale(){
    fetch("/api/v1/sales/products/?from="+this.state.StartDate+"&to="+
      this.state.EndDate+"&type="+this.state.Type+"&export=false").then((res) =>{
        return res.json();
      }).then((res) => {
        this.setState({ProductSales:res});
        console.log(this.state.ProductSales);
      }).catch((err) => {
    });
  }

  exportProductSale(){
    console.log("download");
    window.open("http://localhost:8080/api/v1/sales/products/?from="+this.state.StartDate+"&to="+
      this.state.EndDate+"&type="+this.state.Type+"&export=true")
    /*fetch("/api/v1/sales/products/?from="+this.state.StartDate+"&to="+
      this.state.EndDate+"&type="+this.state.Type+"&export=true",{
        method:"GET",
        headers:{
        'Content-Type': 'text/csv',
        'Accept':'text/csv',
        }
      }).then((res) =>{
        return res.json();
      }).then((res) => {
        console.log(res);
      }).catch((err) => {
    });*/
  }

  changeType(event){
    event.preventDefault()
    console.log();
    this.setState({Type:event.target.value})
  }

  setEndDate(event){
    event.preventDefault()
    this.setState({EndDate:event.target.value})
  }

  setStartDate(event){
    event.preventDefault()
    this.setState({StartDate:event.target.value})
  }

  loadCustomerOrder(){
    if (this.state.CustomerId ==="" ||this.state.CustomerId === 0){
      alert("Please enter a valid customer id")
      return
    }
    fetch("/api/v1/customers/"+this.state.CustomerId+"/orders/").then((res) =>{
        return res.json();
      }).then((res) => {
        this.setState({CustomerOrders:res});
        console.log(this.state.CustomerOrders);
      }).catch((err) => {
    });
  }

  setCustomerId(event){
    event.preventDefault()
    this.setState({CustomerId:event.target.value})
  }

  componentDidMount(){
    fetch("/api/v1/categories/sales/reports/").then((res) => {
      return res.json();
    }).then((res) => {
      this.setState({CategorySale:res});
      console.log(this.state.CategorySale);
    }).catch((err) => {
    });
  }

  generateTimestamp () {
    return moment().format("YYYY-MM-DD h:mm:ss");
  }

  render(){
    return (
    <Container className="box">
      <div>
            <Menu icon size='tiny'>
              <Menu.Item name='side layout' active >
                <Icon name='sidebar' />
              </Menu.Item>
              <Menu.Item name='side layout' active >
                Display Information
              </Menu.Item>
            </Menu>
      </div>
      <Tab menu={{ fluid: true, vertical: true, tabular: true }}
       panes={[
        { menuItem: 'Category Sales', render: () => <Tab.Pane>
          <Table celled>
            <Table.Header>
              <Table.Row>
                <Table.HeaderCell>ID</Table.HeaderCell>
                <Table.HeaderCell>Customer ID</Table.HeaderCell>
                <Table.HeaderCell>Customer FirstName</Table.HeaderCell>
                <Table.HeaderCell>Category ID</Table.HeaderCell>
                <Table.HeaderCell>Category Name</Table.HeaderCell>
                <Table.HeaderCell>Number Purchased</Table.HeaderCell>
              </Table.Row>
            </Table.Header>
            <Table.Body>
            {
              this.state.CategorySale.map((m,index)=>
                {
                  return (
                      <Table.Row key={index+1}>
                          <Table.Cell>{index+1}</Table.Cell>
                          <Table.Cell>{m.customer_id}</Table.Cell>
                          <Table.Cell>{m.customer_first_name}</Table.Cell>
                          <Table.Cell>{m.category_id}</Table.Cell>
                          <Table.Cell>{m.category_name}</Table.Cell>
                          <Table.Cell>{m.number_purchased}</Table.Cell>
                      </Table.Row>
                  )
                })
            }
          </Table.Body>
          </Table>
        </Tab.Pane> },
        { menuItem: 'Product Sales', render: () => <Tab.Pane>
        <Form>
         <Form.Group>
           <Form.Input  width={4}  placeholder="Start Date" type="date" value={this.state.StartDate}
           onChange={this.setStartDate}/>
           <Form.Input  placeholder="End Date"  width={4} type="date" value={this.state.EndDate}
           onChange={this.setEndDate}/>
           <Form.Select width={4} options={[
             { key: 'day', text: 'By Day', value: 'day' },
             { key: 'week', text: 'By Week', value: 'week' },
             { key: 'month', text: 'By Month', value: 'month' },
             { key: 'year', text: 'By Year', value: 'year' },
           ]} placeholder='Type' value={this.state.Type}
           onChange = {(event,{value})=>this.setState({Type:value})}
           />
           <Button  onClick={this.loadProductSale}  width={2}>Load</Button >
           <Button  onClick={this.exportProductSale}  width={2}>Export</Button >
         </Form.Group>
        </Form>
        <Table celled>
          <Table.Header>
            <Table.Row>
              <Table.HeaderCell>ID</Table.HeaderCell>
              <Table.HeaderCell>Date</Table.HeaderCell>
              <Table.HeaderCell>Product ID</Table.HeaderCell>
              <Table.HeaderCell>Product Name</Table.HeaderCell>
              <Table.HeaderCell>Quantity</Table.HeaderCell>
            </Table.Row>
          </Table.Header>
          <Table.Body>
            {
              this.state.ProductSales.map((m,index)=>{
                return (
                  <Table.Row key={index+1}>
                      <Table.Cell>{index+1}</Table.Cell>
                      <Table.Cell>{m.date}</Table.Cell>
                      <Table.Cell>{m.product_id}</Table.Cell>
                      <Table.Cell>{m.product_name}</Table.Cell>
                      <Table.Cell>{m.quantity}</Table.Cell>
                  </Table.Row>
                )
              })
            }
          </Table.Body>
        </Table>
        </Tab.Pane> },
        { menuItem: 'Customers Order', render: () => <Tab.Pane>
          <Grid>
            <Grid.Column width={5}>
              <Input focus placeholder='customer Id' value={this.state.CustomerId}
              onChange={this.setCustomerId}/>
            </Grid.Column>
            <Grid.Column width={3}>
              <Button onClick={this.loadCustomerOrder}>Load</Button>
            </Grid.Column>
          </Grid>
          <Table celled>
            <Table.Header>
              <Table.Row>
                <Table.HeaderCell>ID</Table.HeaderCell>
                <Table.HeaderCell>Order ID</Table.HeaderCell>
                <Table.HeaderCell>Status</Table.HeaderCell>
                <Table.HeaderCell>OrderDate</Table.HeaderCell>
                <Table.HeaderCell>BaseSubtotal</Table.HeaderCell>
                <Table.HeaderCell>Subtotal</Table.HeaderCell>
                <Table.HeaderCell>GrandTotal</Table.HeaderCell>
              </Table.Row>
            </Table.Header>
            <Table.Body>
            {
              this.state.CustomerOrders.map((m)=>{
                return (
                  <Table.Row key={m.id}>
                      <Table.Cell>{m.id}</Table.Cell>
                      <Table.Cell>{m.confirmation_id}</Table.Cell>
                      <Table.Cell>{m.status}</Table.Cell>
                      <Table.Cell>{moment(m.created_at).format("YYYY-MM-DD h:mm:ss")}</Table.Cell>
                      <Table.Cell>{m.base_subtotal}</Table.Cell>
                      <Table.Cell>{m.subtotal}</Table.Cell>
                      <Table.Cell>{m.grand_total}</Table.Cell>
                  </Table.Row>
                )
              })
            }
            </Table.Body>
          </Table>
        </Tab.Pane> },
      ]} />
    </Container>
    );
  }
}
export default Display;
