<h1>Golang Store Inventory Backend</h1>
<p>This is a store inventory backend built with Golang, Makefile, Docker, Migrate, SQLc, Mock, Tests, Gin, Transactions, CRUD and Cron. The backend allows users to add, edit and delete products from a store inventory. The backend also allows users to view order history and invoices.</p>
<h2>Requirements</h2>
<p>To run this backend, you will need the following:</p>
<ul>
<li>Golang</li>
<li>Docker</li>
<li>PostgreSQL</li>
</ul>
<h2>Installation</h2>
<p>To install this backend, you will need to follow these steps:</p>
<ol>
<li>Clone the GitHub repository:</li>
</ol>
<code-block _nghost-ng-c3537242052="" ng-version="0.0.0-PLACEHOLDER"><div _ngcontent-ng-c3537242052="" class="code-block"><!----><pre _ngcontent-ng-c3537242052=""><code _ngcontent-ng-c3537242052="" role="text" class="code-container no-decoration-radius">git clone https://github.com/&lt;your-username&gt;/golang-store-inventory-backend.git
</code></pre><!----></div></code-block>
<ol start="2">
<li>Navigate to the backend directory:</li>
</ol>
<code-block _nghost-ng-c3537242052="" ng-version="0.0.0-PLACEHOLDER"><div _ngcontent-ng-c3537242052="" class="code-block"><!----><pre _ngcontent-ng-c3537242052=""><code _ngcontent-ng-c3537242052="" role="text" class="code-container no-decoration-radius">cd golang-store-inventory-backend
</code></pre><!----></div></code-block>
<ol start="3">
<li>Run the Makefile to build the Docker container:</li>
</ol>
<code-block _nghost-ng-c3537242052="" ng-version="0.0.0-PLACEHOLDER"><div _ngcontent-ng-c3537242052="" class="code-block"><!----><pre _ngcontent-ng-c3537242052=""><code _ngcontent-ng-c3537242052="" role="text" class="code-container no-decoration-radius">make
</code></pre><!----></div></code-block>
<ol start="4">
<li>Start the Docker container:</li>
</ol>
<code-block _nghost-ng-c3537242052="" ng-version="0.0.0-PLACEHOLDER"><div _ngcontent-ng-c3537242052="" class="code-block"><!----><pre _ngcontent-ng-c3537242052=""><code _ngcontent-ng-c3537242052="" role="text" class="code-container no-decoration-radius">docker-compose up
</code></pre><!----></div></code-block>
<ol start="5">
<li>The backend will be available on http://localhost:8080</li>
</ol>
<h2>Usage</h2>
<p>To use this backend, you will need to follow these steps:</p>
<ol>
<li>
<p>Open a web browser and navigate to http://localhost:8080</p>
</li>
<li>
<p>You will see the login screen</p>
</li>
<li>
<p>Login with the following credentials:</p>
<ul>
<li>Username: admin</li>
<li>Password: password</li>
</ul>
</li>
<li>
<p>After logging in, you will see the main backend screen</p>
</li>
<li>
<p>The main screen allows you to add, edit and delete products from a store inventory</p>
</li>
<li>
<p>The main screen also allows you to view order history and invoices</p>
</li>
</ol>
<h2>Makefile</h2>
<p>The Makefile provides an easy way to build, run and test the backend. The Makefile includes the following tasks:</p>
<ul>
<li><code>build</code>: Builds the Docker container</li>
<li><code>run</code>: Starts the Docker container</li>
<li><code>test</code>: Runs the unit tests</li>
<li><code>clean</code>: Removes the Docker container</li>
</ul>
<h2>Docker</h2>
<p>The backend runs in a Docker container. Docker is a platform that makes it easy to package, ship and manage applications in containers. Using Docker allows the backend to run on any platform that supports Docker, such as Linux, macOS and Windows.</p>
<h2>Migrate</h2>
<p>Migrate is a tool that makes it easy to migrate data to a database. Migrate is used to create and apply database migrations to the PostgreSQL database.</p>
<h2>SQLc</h2>
<p>SQLc is a tool that makes it easy to generate SQL code from data models. SQLc is used to generate the SQL code for the backend data model.</p>
<h2>Mock</h2>
<p>Mock is a tool that makes it easy to create mocks for objects. Mocks are used to test the backend code in isolation.</p>
<h2>Tests</h2>
<p>The backend includes unit tests that cover all use cases of the backend. The tests are run using the Makefile.</p>
<h2>Gin</h2>
<p>Gin is a Golang web framework that is used to build the backend. Gin is a fast and efficient framework that is used to build scalable web backends.</p>
<h2>Transactions</h2>
<p>The backend uses transactions to ensure data consistency. Transactions are used to add, edit and delete products from a store inventory.</p>
<h2>CRUD</h2>
<p>The backend provides CRUD operations for products. CRUD operations allow users to add, edit and delete products from a store inventory.</p>
<h2>Cron</h2>
<p>The backend uses Cron to run scheduled tasks. Scheduled tasks are used to update the inventory, send emails and other tasks.</p>
<h2>Contributing</h2>
<p>If you want to contribute to this backend, you can do so by following these steps:</p>
<ol>
<li>Fork the GitHub repository.</li>
<li>Create a new branch for your changes.</li>
<li>Make the necessary changes to the code.</li>
<li>Test your changes.</li>
<li>Commit your changes.</li>
<li>Make a pull request to the main branch.</li>
</ol>
<h2>Thanks</h2>
<p>Thank you for using this backend!</p>