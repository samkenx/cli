/*
Package manifest represents the strucutre of project data and provides behavior
for lifetime management and eased access of data.

goals:
  - the lifetime of the primary type should begin early in the stack
  - where the marshaled data lives should be left to another pkg (named "workspace"?)
  - discuss how to handle marshaled data file watching/reloading
    - this may lead to needing to improve the overall control flow of the app
    - this may lead to guarding all data access for safety and ensuring updated data
  - behavior like constraints should be defined elsewhere and implemented here
    - constraints pkg should not know of this pkg
    - expansion pkg should not know of this pkg
      - setting up expander should occur higher in stack and then passed in
      - secrets should be fed into held expander which deals with secrets itself
*/
package manifest
