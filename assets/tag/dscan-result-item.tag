<dscan-result-item>
    <li>
        <span class="count" style="background-color: { opts.color };">{ opts.count }</span>
        <span class="name">{ opts.name }</span>
    </li>

    <style scoped>
        :scope {
            list-style-type: none;
        }

        li {
            padding: 10px 15px;
            border-top: 1px solid #1c1e22;
        }

        .count {
            background-color: #3a3f44;
            display: block;
            width: 40px;
            height: 40px;
            float: left;
            margin-top: -10px;
            margin-left: -15px;
            margin-bottom: -10px;
            vertical-align: bottom;
            padding: 10px 0;
            text-align: center;
            color: #0c0c0c;
        }

        .name {
            margin-left: 15px;
            text-overflow: ellipsis;
            width: 80%;
            display: inline-block;
            white-space: nowrap;
            overflow: hidden;
            padding: 0 -20px;
            vertical-align: bottom;
        }

    </style>
</dscan-result-item>