import matplotlib.pyplot as plt

filelist = ["051019gray2","051019gray"]

step_name = [
    "user_feature_step",
    "user_fid_step",
    "home_icon_pre_processor",
    "details_page_pre_processor",
    "ad_item_recall",
    "retention_item_classify_step",
    "ad_filter1_step",
    "landing_page_step",
    "pre_ranking_step",
    "user_cross_item_feature_step",
    "cascade_step",
    "item_fid_step",
    "classify_item_step",
    "fill_item_fid_dispatcher_step",
    "predict_dispatcher_step",
    "model_union_dispatcher_step",
    "dh_fid_step",
    "dh_req_total_time_step",
    "dh_ad_models_scores_step",
    "post_ranking_step",
    "ocpx_step",
    "rta_step",
    "show_budget_control_step",
    "ad_filter2_step",
    "activate_force_insert_step",
    "budget_control_step",
    "post_processor_step"
]

resultlist = []

for filename in filelist:
    with open(filename,"r") as f:
        numlist = [0 for i in range(27)]
        count = 0
        line = f.readline()
        while line: 
            split_text = line.split('[')[-1].split(']')[0].split(" ")
            if len(split_text) == 26: 
                for step in split_text:
                    name,name_count = step.split(":")
                    numlist[step_name.index(name)] += int(name_count)
                count += 1
            line = f.readline() 

        resultlist.append([i/count for i in numlist])
    print(count)


print(resultlist[0])
print(sum(resultlist[0]))
print(resultlist[1])
print(sum(resultlist[1]))

x = [i for i in range(1,28)]
plt.figure(figsize=(30,10))

plt.bar([i-0.25 for i in range(1,28)],[i for i in resultlist[0]],width=0.5,label="old")
plt.bar([i+0.25 for i in range(1,28)],[i for i in resultlist[1]],width=0.5,label="new")
plt.xticks(x,labels=[''.join(i.split("_")[:-1]) for i in step_name],rotation=-45,size=15)
plt.yticks([i for i in range(0,55,2)],size=16)
plt.legend(fontsize=26)
plt.savefig("sort.jpg")