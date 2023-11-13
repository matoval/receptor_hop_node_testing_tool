# receptor_hop_node_testing_tool

## Using receptor_testing_tool to run tests

### Verifying that Receptor can accurately send large files via web sockets
* Build binary and add to Receptor image:
    * Build binary: `go build -o receptor_testing_tool main.go`
    * Move bianry to Receptor: `mv ./receptor_testing_tool ../path_to_receptor/packaging/container/`
    * Add bianry to Receptor image by adding this line to Dockerfile: `COPY /receptor_hop_node_testing_tool /usr/bin/receptor_testing_tool`
    * Build Receptor image: `make container`
* Use Receptor image for control node on Openshift
* Run on control node:
    * Create a 1gb file on the control node: `receptor_testing_tool create-file --fileName /tmp/test --fileSize 1000000000`
    * Hash file in order to verify accuracy: `receptor_testing_tool hash-file --fileName /tmp/test`
    * Send file to execution node: `receptor_testing_tool send-file --socket /tmp/control.sock --node test1 --payload /tmp/test --key /etc/receptor/tls/receptor.key --cert /etc/receptor/tls/receptor.crt`
* Run on execution node:
    * work-command config:
    ```yaml
    - work-command:
        workType: echopayload
        command: bash
        params: "-c \"cp /dev/stdin /tmp/test5;\""
    ```
    * Hash file that was sent from the control node in order to verify that both files are identical: `receptor_testing_tool hash-file --fileName /tmp/test5`
