# Item-Service
[![Coverage Status](https://coveralls.io/repos/github/Shangye-space/Item-Service/badge.svg?branch=master)](https://coveralls.io/github/Shangye-space/Item-Service?branch=master)

## How to run service
<span>Make sure you have docker installed.</span>
<a href="https://www.docker.com/get-started">install here!</a>

<p>run commands in following order:</p>
<ul>
    <li><h5>docker-compose build</h5></li>
    <li><h5>docker-compose up</h5></li>
</ul>

### Description
This is one of microservices for "Shangye.space" project. 

<h4>Item service is responsible for:</h4>
<ol>
    <li>API</li>
        <ul>
            <li>Private</li>
            <ul>
                <li>Get items</li>
                <li>Create items</li>
                <li>Edit items</li>
                <li>Update items</li>
                <li>Delete items</li>
            </ul>
            <li>Public</li>
                <ul>
                    <li>Empty for now (TBD)</li>
                </ul>
        </ul>
    <li>Search</li>
        <ul>
            <li>Search engine</li> 
            <ul>
                <li>By name</li> 
                <li>By description</li>
                <li>By category</li>
            </ul>  
        </ul>
    <li>Filters</li>
        <ul>
            <li>By name</li>
            <li>By price</li>
            <li>By manufacturer</li>
        </ul>
    <li>Sharing items in social media</li>
</<li>
