function onclick_btn() {
    axios({
            method: 'get',
            url: "http://mailbook.in.th/api/list",
            responseType: 'application/json'
        }).then(function(response) {
            var mailItems = response.data
            var mailListTable = document.querySelector("#mail-list")

            for (let i = 0; i < mailItems.length; i++) {
                var newTr = document.createElement("tr")
                var newTdPackageId = document.createElement("td")
                var newTdTitle = document.createElement("td")
                var newTdHouseNo = document.createElement("td")
                var newTdDate = document.createElement("td")

                newTdPackageId.innerText = mailItems[i].packageNo
                newTdTitle.innerText = mailItems[i].title
                newTdHouseNo.innerText = mailItems[i].houseNo
                newTdDate.innerText = mailItems[i].createDate

                newTr.appendChild(newTdPackageId)
                newTr.appendChild(newTdTitle)
                newTr.appendChild(newTdHouseNo)
                newTr.appendChild(newTdDate)

                mailListTable.appendChild(newTr)
            }
        })
        .then(function(error) {
            if (error) {
                console.log(error)
            }
        })
        .then(function() {
            console.log("Success fetch api/list")
        })
}