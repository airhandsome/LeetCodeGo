import pandas as pd

# 1.大的国家，选出面基至少为300万或者人口至少为2500万的国家
# 简单的pandas过滤
def big_countries(world: pd.DataFrame) -> pd.DataFrame:
    df = world[(world['area'] >= 3000000) | (world['population'] >= 25000000)]
    return df[['name', 'population', 'area']]

# 1757 可回收且低脂的产品，过滤出可回收且低脂
def find_products(products: pd.DataFrame) -> pd.DataFrame:
    df = products[(products['low_fats'] == 'Y') & (products['recyclable'] == 'Y')]
    return df[['product_id']]

# 183 从不订购的客户，考察且或非中的非，用~表示，而不是用Not表示
def find_customers(customers: pd.DataFrame, orders: pd.DataFrame) -> pd.DataFrame:
    df = customers[~customers['id'].isin(orders['customerId'])]
    df = df[['name']].rename(columns={'name': 'Customers'})
    return df

# 1148 文章浏览1， 考察的不同列相等
def article_views(views: pd.DataFrame) -> pd.DataFrame:
    df = views[views['author_id'] == views['viewer_id']][['author_id']]
    df = df.rename(columns={'author_id': 'id'})
    df = df.drop_duplicates()
    df = df.sort_values(by='id')
    return df

# 1683. 无效的推文，对比不同列的值进行过滤
def article_views(views: pd.DataFrame) -> pd.DataFrame:
    views = views[views['author_id'] == views['viewer_id']]['author_id']
    views = views.rename({'author_id': 'id'})
    views.drop_duplicates(inplace=True)
    views = views.sort_values(by='id')
    return views['id']

# 1683. 无效的推文，列转str，使用str的方法进行过滤
def invalid_tweets(tweets: pd.DataFrame) -> pd.DataFrame:
    return tweets[tweets['content'].str.len() > 15][['tweet_id']]

# 1873. 计算特殊奖金, 通过apply的方式进行复杂的过滤
def calculate_special_bonus(employees: pd.DataFrame) -> pd.DataFrame:
    employees['bonus'] = employees.apply(
        lambda x: x['salary'] if x['employee_id'] % 2 and not x['name'].startswith('M') else 0,
        axis=1
    )
    df = employees[['employee_id', 'bonus']].sort_values('employee_id')
    return df

# 1667. 修复表中的名字 这边注意，使用apply的话，本身就是str，不用强制转
def fix_names(users: pd.DataFrame) -> pd.DataFrame:
    users['name'] = users.apply(
        lambda x: x['name'][0].upper() + x["name"][1:].lower(),
        axis=1
    )
    return users.sort_values('user_id')

# 1517. 查找拥有有效邮箱的用户 正则表达式的应用
def valid_emails(users: pd.DataFrame) -> pd.DataFrame:
    return users[users["mail"].str.match(r"^[a-zA-Z][a-zA-Z0-9_.-]*\@leetcode\.com$")]


# 1527. 患某种疾病的患者 字符串contains正则表达
def find_patients(patients: pd.DataFrame) -> pd.DataFrame:
    return patients[patients['conditions'].str.contains(r'\bDIAB1')]

# 177. 第N高的薪水
def nth_highest_salary(employee: pd.DataFrame, N: int) -> pd.DataFrame:
    df = employee[['salary']].drop_duplicates()
    key = 'getNthHighestSalary(' + str(N) + ')'
    if len(df) < N:
        #这边不能直接返回None
        return pd.DataFrame({key: [None]})

    # 这边类似数据库的count N limit 1
    df = df.sort_values(by='salary', ascending=False).head(N).tail(1)
    return df.rename(columns={'salary': key})

# 176. 第二高的薪水
def second_highest_salary(employee: pd.DataFrame) -> pd.DataFrame:
    df = employee[['salary']].drop_duplicates()
    if len(df) == 1:
        return pd.DataFrame({'SecondHighestSalary': None})

    df = df.sort_values(by='salary', ascending=False).head(2).tail(1)
    # 这边就换种写法把，不用rename了
    return pd.DataFrame({'SecondHighestSalary': df['salary'][0]})


def department_highest_salary(employee: pd.DataFrame, department: pd.DataFrame) -> pd.DataFrame:
    # 类似sql的left join on a.departmentId = b.id
    df = employee.merge(department, left_on='departmentId', right_on='id', how='left')
    df.rename(columns={'name_x': 'Employee', 'name_y': 'Department', 'salary': 'Salary'}, inplace=True)

    # 先找出最高工资，groupby分组, transform是执行某个函数，可以用lambda自己定义，也可以输入约定速成的几个
    max_salary = df.groupby('Department')['Salary'].transform('max')
    # 过滤工资等于最高工资的员工
    df = df[df['Salary'] == max_salary]

    return df[['Department', 'Employee', 'Salary']]

# 178. 分数排名
def order_scores(scores: pd.DataFrame) -> pd.DataFrame:
    # 类似sql的 dense_rank over (partition score order by score desk)
    scores['rank'] = scores['score'].rank(method='dense', ascending=False)
    return scores[['score', 'rank']].sort_values('score', ascending=False)


# 196. 删除重复的电子邮箱
def delete_duplicate_emails(person: pd.DataFrame) -> None:
    person.sort_values('id', inplace=True)
    person.drop_duplicates(['email'], keep='first', inplace=True)
    return

# 1795. 每个产品在不同商店的价格
def rearrange_products_table(products: pd.DataFrame) -> pd.DataFrame:
    df =products.set_index(['product_id']).stack(dropna=True)
    df = df.reset_index()
    df.columns = ['product_id', 'store', 'price']
    return df

# 1907. 按分类统计薪水
def count_salary_categories(accounts: pd.DataFrame) -> pd.DataFrame:
    low = (accounts['income'] < 20000).sum()
    avg = ((accounts['income'] <= 50000) & (accounts['income'] >= 20000)).sum()
    high = (accounts['income'] > 50000).sum()
    return pd.DataFrame({
        'category': ['Low Salary', 'Average Salary', 'High Salary'],
        'accounts_count': [low, avg, high]
    })

# 1741. 查找每个员工花费的总时间
def total_time(employees: pd.DataFrame) -> pd.DataFrame:
    df = employees.groupby(by=['event_day', 'emp_id']).sum().reset_index().rename(columns={'event_day': 'day'})
    df['total_time'] = df['out_time'] - df['in_time']
    df.drop(columns=['in_time', 'out_time'], inplace=True)
    return df.sort_values(by='emp_id')

# 511. 游戏玩法分析 I
def game_analysis(activity: pd.DataFrame) -> pd.DataFrame:
    df = activity.sort_values(by=['player_id', 'event_date'])[['player_id', 'event_date']].rename(columns={'event_date': 'first_login'})
    print(df)
    return df.drop_duplicates('player_id')

# 2356. 每位教师所教授的科目种类的数量
def count_unique_subjects(teacher: pd.DataFrame) -> pd.DataFrame:
    df = teacher.drop_duplicates(['teacher_id', 'subject_id'])
    df['cnt'] = df.groupby('teacher_id')['subject_id'].transform('count')
    return df[['teacher_id', 'cnt']].drop_duplicates(['teacher_id', 'cnt'])

# 596. 超过5名学生的课
def find_classes(courses: pd.DataFrame) -> pd.DataFrame:
    courses['cnt'] = courses.groupby('class')['student'].transform('count')
    df = courses[courses['cnt'] >= 5].drop('student', axis=1)[['class']]
    return df.drop_duplicates()

# 586. 订单最多的客户
def largest_orders(orders: pd.DataFrame) -> pd.DataFrame:
    if orders.empty:
        return pd.DataFrame({'customer_number': []})

    df = orders.groupby('customer_number').size().reset_index(name='count')
    df.sort_values(by='count', ascending=False, inplace=True)
    return df[['customer_number']][:1]

# 1484. 按日期分组销售产品
def categorize_products(activities: pd.DataFrame) -> pd.DataFrame:
    groups = activities.groupby('sell_date')
    stats = groups.agg(
        num_sold=('product', 'nunique'),
        products=('product', lambda x: ','.join(sorted(set(x))))
        ).reset_index()

    stats.sort_values('sell_date', inplace=True)
    return stats

# 1693. 每天的领导和合伙人
def daily_leads_and_partners(daily_sales: pd.DataFrame) -> pd.DataFrame:
    df = daily_sales.groupby(['date_id', 'make_name']).nunique().reset_index()
    df = df.rename(columns={'lead_id':'unique_leads', 'partner_id': 'unique_partners'})
    return df

# 1050. 合作过至少三次的演员和导演
def actors_and_directors(actor_director: pd.DataFrame) -> pd.DataFrame:
    df = actor_director.groupby(['actor_id', 'director_id']).size().reset_index(name='cnt')
    df = df[df['cnt'] >= 3]
    return df[['actor_id', 'director_id']]

# 1378. 使用唯一标识码替换员工ID
def replace_employee_id(employees: pd.DataFrame, employee_uni: pd.DataFrame) -> pd.DataFrame:
     df = employees.merge(employee_uni, how='left', left_on='id', right_on='id')
     return df[['unique_id', 'name']]

# 1280. 学生们参加各科测试的次数
def students_and_examinations(students: pd.DataFrame, subjects: pd.DataFrame,
                              examinations: pd.DataFrame) -> pd.DataFrame:
    # 按 id 和科目分组，并计算考试次数。
    grouped = examinations.groupby(['student_id', 'subject_name']).size().reset_index(name='attended_exams')
    all_id_subjects = pd.merge(students, subjects, how='cross')
    id_subjects_count = pd.merge(all_id_subjects, grouped, on=['student_id', 'subject_name'], how='left')
    id_subjects_count['attended_exams'] = id_subjects_count['attended_exams'].fillna(0).astype(int)
    id_subjects_count.sort_values(['student_id', 'subject_name'], inplace=True)

    return id_subjects_count[['student_id', 'student_name', 'subject_name', 'attended_exams']]

# 570. 至少有5名直接下属的经理
def find_managers(employee: pd.DataFrame) -> pd.DataFrame:
    df = employee.groupby(['managerId']).size().reset_index(name='cnt')
    df = df[df['cnt'] >= 5]
    df = df.merge(employee, how='inner', right_on='id', left_on='managerId')
    return df[['name']]

# 607. 销售员
def sales_person(sales_person: pd.DataFrame, company: pd.DataFrame, orders: pd.DataFrame) -> pd.DataFrame:
    company = company[company['name'] == 'RED']
    df = orders.merge(company, how='inner', on='com_id')
    sales_person = sales_person[~sales_person['sales_id'].isin(df['sales_id'])]
    return sales_person[['name']]