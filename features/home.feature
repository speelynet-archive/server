Feature: the home page route
  In order to display the home page
  As a web server
  I need to properly route the files

  Background:
    Given I create a mock server

  @wip
  Scenario: index.html
    Given I create the temporary file "static/index.html"
    And the content of the temporary file is
      """
      <html>
        Hello, World!
      </html>
      """
    When I go to "mock.server/"
    Then I should see the file content

  @wip
  Scenario: other files
    Given I create the temporary file "static/thing.txt"
    And the content of the temporary file is
        """
        Hello, World!
        """
    When I go to "mock.server/thing.txt"
    Then I should see the file content
