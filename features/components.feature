Feature: the components route
  In order to provide access to the component library
  As a web server
  I need to route the directories properly

  Background:
    Given I create a mock server

  Scenario Template: stable directory
    Given I create the temporary file "components/stable/index.js"
    And the content of the temporary file is
    """
    function test() {
      console.log("Hello, World!");
    }
    """
    When I go to "mock.server/components/<path>index.js"
    Then I should see the file content

    Scenarios:
      | path    |
      |         |
      | stable/ |

  Scenario: latest directory
    Given I create the temporary file "components/latest/index.js"
    And the content of the temporary file is
    """
    function test() {
      console.log("foobar");
    }
    """
    When I go to "mock.server/components/latest/index.js"
    Then I should see the file content

  Scenario Template: root is index.js
    Given I create the temporary file "components/<path>/index.js"
    And the content of the temporary file is
    """
    function test() {
      console.log("<url>");
    }
    """
    When I go to "mock.server/components<url>"
    Then I should see the file content

    Scenarios:
      | path   | url      |
      | stable |          |
      | stable | /        |
      | stable | /stable  |
      | stable | /stable/ |
      | latest | /latest  |
      | latest | /latest/ |
